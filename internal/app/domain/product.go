package domain

type Product struct {
	name          string
	price         *Price
	weightInGrams int
}

func (p *Product) GetPrice() *Price {
	return p.price
}

func (p *Product) GetProductName() string {
	return p.name
}

func (p *Product) GetWeightInGrams() int {
	return p.weightInGrams
}

func NewProduct(name string, price *Price, weightInGrams int) Product {
	return Product{
		name:          name,
		price:         price,
		weightInGrams: weightInGrams,
	}
}
