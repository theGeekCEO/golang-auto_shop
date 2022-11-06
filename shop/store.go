package shop

import (
	"errors"
	"fmt"
)

type Store struct {
	ID     string
	Stock  map[string]Product // [itemId]Product
	Orders []Order
}

func (s *Store) AddItem(item ProductItem, price float64) {
	// check if there are existing product with that item
	itemId := item.GetID()
	if p, e := s.Stock[itemId]; e {
		p.QtyInStock++  // increment quantity
		p.Price = price // update price if needed
		s.Stock[itemId] = p
	} else {
		s.Stock[itemId] = Product{
			ID:         GenerateID(),
			Item:       item,
			QtyInStock: 1,
			Price:      price,
		}
	}
}

func (s *Store) GetItems() []ProductItem {
	var items []ProductItem
	for _, v := range s.Stock {
		items = append(items, v.Item)
	}

	return items
}

func (s *Store) DisplayItems() {
	for i, item := range s.GetItems() {
		fmt.Printf("%d: %s => %s\n", i+1, item.GetID(), item.GetName())
	}
}

func (s *Store) SellItem(item ProductItem) (bool, error) {
	itemId := item.GetID()

	if p, e := s.Stock[itemId]; e {
		if !p.IsAvailable() {
			return false, errors.New("product not in stock")
		}

		if p.QtyInStock == 1 {
			// remove item from stock completely
			delete(s.Stock, itemId)
		} else {
			p.QtyInStock--
			s.Stock[itemId] = p
		}

		// add item to orders
		s.Orders = append(s.Orders, Order{&p, p.Price})
		return true, nil
	}

	return false, errors.New("product not in stock")
}

// GetOutflow - Get's list of sold items and total price
func (s *Store) GetOutflow() ([]Order, float64) {
	totalPrice := float64(0)
	for _, o := range s.Orders {
		totalPrice += o.SoldPrice
	}

	return s.Orders, totalPrice
}

// DisplayOutflow - Display's list of sold items and total price
func (s *Store) DisplayOutflow() {
	fmt.Println("List of sold items and total price")

	orders, totalPrice := s.GetOutflow()
	for i, o := range orders {
		fmt.Printf("%d: %s => %f\n", i+1, o.Product.Item.GetName(), o.SoldPrice)
	}

	fmt.Printf("Total Price is: %f\n", totalPrice)
}
