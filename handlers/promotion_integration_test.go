//go:build integration

package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	handler "unittest/handlers"
	"unittest/repositories"
	"unittest/service"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {

	amount := 100
	expected := 80

	// Mock data จาก Repo
	promoRepo := repositories.NewPromotionRepository_Mock()
	// GetPromotion มาจาก Repo mock data
	promoRepo.On("GetPromotion").Return(repositories.Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 20,
	}, nil)

	// Func ของ Service
	promoService := service.NewPromotionServie(promoRepo)

	// Func ของ Handler
	promoHandler := handler.NewPromotionHandler(promoService)



	//******** ทำ fiber *******
	
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		// น่าจะแทน postman ในการส่งคำสั่ง
		// method, target(path),body
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		// เช็ค status code ก่อน ว่า 200 หรือไม่ ถ้าไม่ใช้ 200 ไม่ต้องอ่าน Body
		ok := assert.Equal(t, fiber.StatusOK, res.StatusCode)
		if ok {
			// อ่านค่าจาก Body เพราะต้องมาเช็คค่า expected ว่าตรงตามต้องการรึป่าว
			body, _ := io.ReadAll(res.Body)
			fmt.Println("***************")
			fmt.Println(string(body))
			fmt.Println("***************")
			assert.Equal(t, strconv.Itoa(expected), string(body)) // OK
		}

}
