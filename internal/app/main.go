package main

import (
	"github.com/hiteshpattanayak-tw/DDD/internal/app/model"
)

func main() {
	ipad := model.NewProduct("IPad Air 2020", model.NewPrice("INR", 60000))
	pen := model.NewProduct("Hero ink pen", model.NewPrice("INR", 10))
	bat := model.NewProduct("GM Cricket bat", model.NewPrice("INR", 2000))

	item1 := model.NewItem(ipad, 1)
	item2 := model.NewItem(pen, 1)
	item3 := model.NewItem(bat, 2)

	cart := model.NewCart()
	cart.Add(item1)
	cart.Add(item2)
	cart.Add(item3)
}
