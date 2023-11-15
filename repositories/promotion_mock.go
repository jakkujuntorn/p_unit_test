package repositories

import (
	"github.com/stretchr/testify/mock"
)

type promotionRepository_Mock struct {
	// ทำ mock data
	mock.Mock
}

// mock data ต้อง return เป็น pointer เท่านั้น ไม่งั้น error ******
func NewPromotionRepository_Mock() *promotionRepository_Mock {
	return &promotionRepository_Mock{}
}

// อย่างน้อยต้องมี interface เพื่อเอามาทำ test func
func (m *promotionRepository_Mock) GetPromotion() (Promotion, error) {
	// GetPromotion ไม่รับ paramiter  m.Called เลยไม่ต้องส่งอะไรเข้าไป
	args := m.Called()

	// case typeตอน return เพราะ func return Promotion
	return args.Get(0).(Promotion), args.Error(1)
}
