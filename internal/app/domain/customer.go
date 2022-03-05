package domain

type Customer struct {
	address *Address
	accounts []*Account
}

func NewCustomer(address *Address) *Customer {
	return &Customer{
		address:  address,
		accounts: make([]*Account, 0),
	}
}

func (c *Customer) AddAccount(account *Account) {
	c.accounts = append(c.accounts, account)
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
