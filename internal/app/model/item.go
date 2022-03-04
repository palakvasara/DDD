package model

type Item struct {
	Product Product
	Qty     int
}

func NewItem(product Product, qty int) Item {
	return Item{
		Product: product,
		Qty:     qty,
	}
}

func (i Item) GetProductName() string {
	return i.Product.Name
}
