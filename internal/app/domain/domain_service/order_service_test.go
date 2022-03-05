package domain_service_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/domain"
	"github.com/hiteshpattanayak-tw/DDD/internal/app/domain/domain_service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddCart_Checkout(t *testing.T) {
	product := domain.NewProduct("IPad", domain.NewPrice("INR", 60000))
	item := domain.NewItem(product, 2)

	cart := domain.NewCart()
	cart.Add(item)

	products := []domain.Product{
		product,product,
	}
	expected := domain.NewOrder(products)

	os := domain_service.NewOrderService()
	actual := os.Checkout(cart)
	assert.Equal(t, expected, actual)
	assert.True(t, cart.IsCheckedOut())
	assert.Empty(t, cart.Items)
}
