package model

type Product struct {
	Name string
}

func NewProduct(name string) Product {
	return Product{
		Name: name,
	}
}
