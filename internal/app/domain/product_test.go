package domain_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/domain"
	"github.com/hiteshpattanayak-tw/DDD/internal/app/domain/domain_service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct_ShouldReturnNewProduct(t *testing.T) {
	name := "pen"
	currency := "INR"
	value := 20.1
	price := domain.NewPrice(currency, value)

	domain_service.AddNewProductToCompetitorsList(name, *price)
	discountedPrice := domain_service.GetDiscountedPrice(name)

	expected := domain.Product{Name: name, Price: &discountedPrice}
	actual := domain.NewProduct(name, &discountedPrice)

	assert.Equal(t, expected, actual)
}
