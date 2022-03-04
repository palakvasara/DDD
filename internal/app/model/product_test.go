package model_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct_ShouldReturnNewProduct(t *testing.T) {
	name := "pen"
	expected := model.Product{Name: name}
	actual := model.NewProduct(name)

	assert.Equal(t, expected, actual)
}
