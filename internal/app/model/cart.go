package model

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Cart struct {
	id    string
	Items []Item
	DomainEvents []string
}

func NewCart() *Cart {
	id, _ := gonanoid.New()
	return &Cart{
		id: id,
	}
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

		c.Items = append(c.Items[0:idx], c.Items[idx+1:len(c.Items)]...)

		c.apply(i, "removed")
	}
}

func (c *Cart) IsEqualTo(anotherC *Cart) bool {
	return c.id == anotherC.id
}

func (c *Cart) apply(item Item, action string) {
	c.DomainEvents = append(c.DomainEvents, fmt.Sprintf("%s was %s", item.GetProductName(), action))
}
