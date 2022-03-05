package domain

type Order struct {
	products []Product
}

func NewOrder(products []Product) *Order {
	return &Order{products: products}
}
