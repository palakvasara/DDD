package model

import "fmt"

type Cart struct {
	Items []Item
	DomainEvents []string
}

func NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) Add(item Item) {
	c.Items = append(c.Items, item)
	c.apply(item, "added")
}

func (c *Cart) Remove(item Item) {
	for idx, i := range c.Items {
		if i.GetProductName() != item.GetProductName() {
			continue
		}

		c.Items = append(c.Items[0: idx], c.Items[idx+1: len(c.Items)]...)

		c.apply(i, "removed")
	}
}

func (c *Cart) apply(item Item, action string) {
	c.DomainEvents = append(c.DomainEvents, fmt.Sprintf("%s was %s", item.GetProductName(), action))
}