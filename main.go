package main

import (
	"fmt"

	"github.com/sunday4me/Backend_Golang_S_S_E_Project/shop"
)

func main() {
	item1 := shop.NewCar("Lexus", "Camry", "red")
	item2 := shop.NewCar("Lexus", "Camry", "blue")
	item3 := shop.NewCar("Mutang", "Golf", "orange")

	store := shop.NewStore()
	store.AddItem(item1, 30000)
	store.AddItem(item1, 30000)
	store.AddItem(item1, 30000)

	store.AddItem(item2, 40000)
	store.AddItem(item2, 40000)

	store.AddItem(item3, 50000)

	fmt.Println("Initial products")
	fmt.Println()
	store.DisplayItems()

	fmt.Println("Selling some items")
	fmt.Println()
	ok, err := store.SellItem(item1)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
	}

	if ok {
		fmt.Println("Sold =>", item1.GetName())
		fmt.Println()
	}

	store.DisplayOutflow()

	fmt.Println("Sell another item")
	fmt.Println()
	ok, err = store.SellItem(item2)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
	}

	if ok {
		fmt.Println("Sold =>", item1.GetName())
		fmt.Println()
	}

	store.DisplayOutflow()

	fmt.Println("Done with preview. Thanks for previewing")
}
