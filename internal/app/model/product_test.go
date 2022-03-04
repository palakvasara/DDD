package model_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct_ShouldReturnNewProduct(t *testing.T) {
	name := "pen"
	price := &model.Price{
		Currency: "INR",
		Value:    10.1,
	}
	expected := model.Product{Name: name, Price: price}
	actual := model.NewProduct(name, price)

	assert.Equal(t, expected, actual)
}
