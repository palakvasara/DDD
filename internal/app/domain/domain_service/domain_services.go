package domain_service

import "github.com/hiteshpattanayak-tw/DDD/internal/app/domain"

type CompetitorsBasedPricerService interface {
	AddNewProductToCompetitorsList(product string, price domain.Price)
	GetDiscountedPrice(productName string) domain.Price
}
