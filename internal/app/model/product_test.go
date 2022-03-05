package model_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model/domain_service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct_ShouldReturnNewProduct(t *testing.T) {
	name := "pen"
	currency := "INR"
	value := 20.1
	price := model.NewPrice(currency, value)

	domain_service.AddNewProductToCompetitorsList(name, *price)
	discountedPrice := domain_service.GetDiscountedPrice(name)

	expected := model.Product{Name: name, Price: &discountedPrice}
	actual := model.NewProduct(name, &discountedPrice)

	assert.Equal(t, expected, actual)
}
