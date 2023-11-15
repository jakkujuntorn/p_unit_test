package service_test

import (
	"errors"
	"testing"
	"unittest/repositories"
	"unittest/service"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)


// การ TEst ตรงนี้ จาก service ไปหา repository
func Test_PromotionDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	// ทำ test case ให้หลากหลาย
	cases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
		// {name: "not applied zero", purchaseMin: 100, discountPercent: 20, amount: 0, expected: 0},
	}

	for _, v := range cases {
		// **** ทำ sub test ****
		t.Run(v.name, func(t *testing.T) {

			//******** 1.Mock data ***************
			// 1.Arrage เตรียมของ *****************
			// ทำ unit test ต้องใช้  mock data  เขาไม่ใช้ data จริงทำกัน
			// ที่ต้องทำตรงนี้เพราะ NewPromotionServie ต้องการ I_PromotionRepository
			//ใช้ mock data จาก layer ที่ต่ำกว่า เพราะเราต้องเอาไปต่อกับ Func ของ Service / repo
			promoRepo := repositories.NewPromotionRepository_Mock()

			//******** 2. Mock GetPromotion ******
			// GetPromotion มาจาก repositories
			// GetPromotion จะ return Promotion กับ error
			// promoRepo ทำ recive แบบ test  ให้เหมือน I_PromotionRepository
			// เอา struct ของ mock มารวมกับ test case ที่ ทำขึ้น
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     v.purchaseMin,
				DiscountPercent: v.discountPercent,
			}, nil)

			//*******************  เริ่ม Test ****************
			// ส่ง promoRepo ที่ทำ recive ให้ conform ตาม I_PromotionRepository
			promoService := service.NewPromotionServie(promoRepo)

			// ******* 2. Test Func *******
			// 2.Act **********************
			// Func ที่ต้องการเทส คือ CalculateDiscount
			// ตรงนี้ไม่มี error เพราะเทสแบบ success
			// CalculateDiscount เป็น func ที่ทำงานจริง มี code การทำงานจริง
			discount, _ := promoService.CalculateDiscount(v.amount)
			expected := v.expected

			// 3.Assert ********************
			// lib ที่ใช้ test
			assert.Equal(t, expected, discount)

			

		})

	}

	// ************** Error ZeroAmount *************
	// **** test เฉพาะบางเงื่อนไขใน Func ใช้ AssertNotCalled ****
	t.Run("Zero amount", func(t *testing.T) {
		
		//********* Mock data *********
		//ใช้ mock data จาก layer ที่ต่ำกว่า เพราะเราต้องเอาไปต่อกับ Func ของ Service / repo
		promoRepo := repositories.NewPromotionRepository_Mock()

		// Arrage
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := service.NewPromotionServie(promoRepo)

		//Act
		// ส่ง 0 เข้าไป เพื่อให้ return error ออกมา
		_, err := promoService.CalculateDiscount(0)

		// 3.Assert
		// ******test error ว่าเป็น error ที่เราเซตไว้รึป่าว *****
		// error ตัวนี้มาจาก CalculateDiscount / amount <= 0
		// error ที่ได้ ต้องเป็น ErrZeroAmount เท่านั้น
		// ใช้ ErrorIs เพราะ ต้องการ error ที่เป็นของเราเองด้วย
		// พารามิเตอร์ ( t, err, error ที่จะเช็คว่าตรงกันไหม )
		assert.ErrorIs(t, err, service.ErrZeroAmount)

		//***********  สั่งห้ามไป call GetPromotion() **********
		// กรณี Code เช็คค่า amount อยู่ด้านล่าง Code GetPromotion() จะ มี error ออกมา
		promoRepo.AssertNotCalled(t, "GetPromotion")

		// promoRepo.AssertNotCalled(t, "TestCalled")
	})

	// ************* Error Repository ******************
	t.Run("Repository Error", func(t *testing.T) {
		//ใช้ mock data จาก layer ที่ต่ำกว่า เพราะเราต้องเอาไปต่อกับ Func ของ Service / repo
		promoRepo := repositories.NewPromotionRepository_Mock()

		//******** Mock data *********
		// Arrage
		// ตรงนี้ไม่ต้องส่ง mock data แต่ส่ง error เข้าไปแทน เพราะจะเทส error จากตรงนี้
		// สร้าง error ให้ return ออกมา
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("Erro Naja"))

		promoService := service.NewPromotionServie(promoRepo)
 
		//Act
		// ส่ง 100 เข้าไป เพื่อให้ return error ออกมา
		// error ตัวนี้มาจาก CalculateDiscount / GetPromotion
		// การส่ง100 เข้าไปจะผ่าน การเช็ค error เรื่องจำนวนเงิน
		_, err := promoService.CalculateDiscount(100)

		//Assert
		//***********  สั่งห้ามไป call GetPromotion() **********
		// error ที่ได้ ต้องเป็น ErrRepository เท่านั้น
		assert.ErrorIs(t, err, service.ErrRepository)
		// assert.ErrorIs(t, err,err)

		
	})

}
