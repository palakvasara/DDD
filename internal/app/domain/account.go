package domain

type Account struct {
	address *Address
}

func NewAccount(address *Address) *Account {
	return &Account{address: address}
}

func (a *Account) UpdateAddresses(address *Address) {
	a.address = address
}

func (a *Account) GetAddress() *Address {
	return a.address
}