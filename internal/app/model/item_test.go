package model_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewItem_ShouldReturnNewItem(t *testing.T) {
	product := model.NewProduct("IPad", model.NewPrice("INR", 60000))
	expected := model.Item{Product: product, Qty: 1}
	actual := model.NewItem(product, 1)

	assert.Equal(t, expected, actual)
}
