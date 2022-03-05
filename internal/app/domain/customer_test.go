package domain_test

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomer_ShouldUpdateAddress(t *testing.T) {
	cityName := "Pune"
	address := domain.NewAddress(cityName)

	account1 := domain.NewAccount(address)
	account2 := domain.NewAccount(address)

	customer := domain.NewCustomer(address)
	customer.AddAccount(account1)
	customer.AddAccount(account2)

	cityName = "Mumbai"
	newAddress := domain.NewAddress(cityName)
	customer.UpdateAddresses(newAddress)

	assert.Equal(t, newAddress, customer.GetAddress())
	assert.Equal(t, newAddress, account1.GetAddress())
	assert.Equal(t, newAddress, account2.GetAddress())
}