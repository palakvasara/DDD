package model

type Price struct {
	Currency string
	Value    float64
}

func NewPrice(currency string, value float64) *Price {
	return &Price{
		Currency: currency,
		Value:    value,
	}
}
