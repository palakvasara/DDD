package model_test

import (
	"fmt"
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddCart_ShouldAddNewProductToCart(t *testing.T) {
	product := model.NewProduct("IPad")
	item := model.NewItem(product, 2)

	expected := &model.Cart{Items: []model.Item{item}}

	cart := model.NewCart()
	cart.Add(item)

	assert.ElementsMatch(t, expected.Items, cart.Items)
}

func TestRemoveCart_ShouldRemoveProductFromCartAndAddItToRemovedItems(t *testing.T) {
	product1 := model.NewProduct("IPad")
	product2 := model.NewProduct("Pen")
	item1 := model.NewItem(product1, 2)
	item2 := model.NewItem(product2, 1)

	domainEvents := []string {
		fmt.Sprintf("%s was added", item1.GetProductName()),
		fmt.Sprintf("%s was added", item2.GetProductName()),
		fmt.Sprintf("%s was removed", item2.GetProductName()),
	}

	expected := &model.Cart{Items: []model.Item{item1}, DomainEvents: domainEvents}

	cart := model.NewCart()
	cart.Add(item1)
	cart.Add(item2)

	cart.Remove(item2)

	assert.ElementsMatch(t, expected.Items, cart.Items)
	assert.ElementsMatch(t, expected.DomainEvents, cart.DomainEvents)
}

func TestIsEqualTo_ShouldReturnFalseForDifferentCarts(t *testing.T) {
	product := model.NewProduct("IPad")
	item := model.NewItem(product, 2)

	cart1 := model.NewCart()
	cart2 := model.NewCart()

	cart1.Add(item)
	cart2.Add(item)

	assert.False(t, cart1.IsEqualTo(cart2))
}

func TestIsEqualTo_ShouldReturnTrueForSameCarts(t *testing.T) {
	product := model.NewProduct("IPad")
	item := model.NewItem(product, 2)

	cart1 := model.NewCart()
	cart1.Add(item)

	assert.True(t, cart1.IsEqualTo(cart1))
}
