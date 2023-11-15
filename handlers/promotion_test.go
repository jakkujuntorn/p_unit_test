package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	handler "unittest/handlers"
	"unittest/service"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_PromotionCalculateDiscount(t *testing.T) {
	//************ Mock Service ********
	t.Run("Success", func(t *testing.T) {

		// Arrange
		amont := 100
		expected := 80

		// Mock data
		//ใช้ mock data จาก layer ที่ต่ำกว่า เพราะเราต้องเอาไปต่อกับ Func ของ Handler / service ****
		promoService := service.NewPromotionService_Mock()
		//CalculateDiscount มีการรับ paramiter .On เลยต้องส่ง paramiter เข้าไป
		promoService.On("CalculateDiscount", amont).Return(expected, nil)

		// implement
		// promoService conform ตาม interface I_PromotionService
		promoHandler := handler.NewPromotionHandler(promoService)

		// ทำ fiber handler
		// http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		// น่าจะแทน postman ในการส่งคำสั่ง
		// method, target(path),body( body  ไม่ได้ส่ง เป็น nil)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amont), nil)

		// ใน vdo อื่นใช้ มัน return ค่าอะไรออกมา
		// rec:= httptest.NewRecorder()

		// Act
		// fiber มี func การเทสด้วย
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
			//expected กับ body ยังไงมันก็ค่าเดียวกัน มะนจะไม่ตรงกันได้ยังไง ******
			assert.Equal(t, strconv.Itoa(expected), string(body)) // OK
			// assert.Equal(t, "50", string(body)) // FAIL
		}

	})

}
