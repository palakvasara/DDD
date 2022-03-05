package domain_service

import "github.com/hiteshpattanayak-tw/DDD/internal/app/domain"

const discount = 0.1

var CompetitorsProductList map[string]domain.Price

func init() {
	CompetitorsProductList = make(map[string]domain.Price)
}

func AddNewProductToCompetitorsList(product string, price domain.Price) {
	CompetitorsProductList[product] = price
}

func GetDiscountedPrice(productName string) domain.Price {
	for prodName, price := range CompetitorsProductList {
		if prodName != productName {
			continue
		}
		val := price.GetValue()
		price := domain.NewPrice("INR", val * (1- discount))
		return *price
	}

	return domain.Price{}
}
