package handler

import (
	"strconv"
	"unittest/service"

	"github.com/gofiber/fiber/v2"
)

type I_PromotiomHandler interface {
	CalculateDiscount(*fiber.Ctx) error
}

type promotionHandler struct {
	promoService service.I_PromotionService
}

func NewPromotionHandler(promo_Service service.I_PromotionService) I_PromotiomHandler {
	return &promotionHandler{promoService: promo_Service}
}

func (p *promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	// ในนี้มี 2 หัวข้อที่ต้อง เทส

	// 1
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// 2 check discount
	discount, err := p.promoService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendString(strconv.Itoa(discount))
}
