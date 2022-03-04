package model_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddCart_ShouldReturnNewPrice(t *testing.T) {
	expected := &model.Price{
		Currency: "INR",
		Value:    20.1,
	}

	actual := model.NewPrice("INR", 20.1)

	assert.Equal(t, expected, actual)
}
