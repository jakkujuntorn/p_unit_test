package service

import (
	_ "fmt"
	"unittest/repositories"
)

type I_PromotionService interface {
	CalculateDiscount(int) (int, error)
}

type promotionService struct {
	promoRepo repositories.I_PromotionRepository
}

func NewPromotionServie(promo_Repo repositories.I_PromotionRepository) I_PromotionService {
	return &promotionService{
		promoRepo: promo_Repo,
	}
}

func (p *promotionService) CalculateDiscount(amount int) (int, error) {
	// TestCalled()

	// Zero amount
	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	// Error Repository
	// ดึงราคาจาก mock data
	promo, err := p.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	// if amount <= 0 {
	// 	return 0, ErrZeroAmount
	// }

	// ซื้อเกินราคา promotion
	if amount >= promo.PurchaseMin {
		return amount - (promo.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}

// func TestCalled()  {
// 	fmt.Println("Call back")
// }
