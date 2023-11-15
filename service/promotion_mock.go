package service

import (
	"github.com/stretchr/testify/mock"
)

type promotionService_Mock struct {
	mock.Mock
}

func NewPromotionService_Mock() *promotionService_Mock {
	return &promotionService_Mock{}
}

// conform ตาม interface I_PromotionService
func (p *promotionService_Mock) CalculateDiscount(amount int) (int, error) {
	// CalculateDiscount มีการรับ amount เข้ามา Called เลยต้องส่ง amount เข้าไปด้วย
	args := p.Called(amount)

	return args.Int(0), args.Error(1)
}
