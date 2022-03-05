package domain_service

import "github.com/hiteshpattanayak-tw/DDD/internal/app/model"

const discount = 0.1

var CompetitorsProductList map[string]model.Price

func init() {
	CompetitorsProductList = make(map[string]model.Price)
}

func AddNewProductToCompetitorsList(product string, price model.Price) {
	CompetitorsProductList[product] = price
}

func GetDiscountedPrice(productName string) model.Price {
	for prodName, price := range CompetitorsProductList {
		if prodName != productName {
			continue
		}
		val := price.GetValue()
		price := model.NewPrice("INR", val * (1- discount))
		return *price
	}

	return model.Price{}
}
