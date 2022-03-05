package domain

type Customer struct {
	address *Address
	accounts []*Account
}

func NewCustomer(address *Address, accounts []*Account) *Customer {
	return &Customer{
		address:  address,
		accounts: accounts,
	}
}

func (c *Customer) UpdateAddresses(address *Address) {
	c.address = address
	for _, ac := range c.accounts {
		ac.UpdateAddresses(address)
	}
}

func (c *Customer) GetAddress() *Address {
	return c.address
}
