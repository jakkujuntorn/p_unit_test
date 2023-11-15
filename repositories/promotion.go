package repositories

import (
	
)

type Promotion struct {
	ID              int
	PurchaseMin     int
	DiscountPercent int
}

type I_PromotionRepository interface {
	GetPromotion() (Promotion, error)
}
