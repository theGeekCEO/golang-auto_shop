package shop

import "fmt"

type ProductItem interface {
	Details()
	GetID() string
	GetName() string
}

type Product struct {
	ID         string
	Item       ProductItem
	Price      float64
	QtyInStock int
}

func (p *Product) Display() {
	fmt.Printf("Product ID: %s\nPrice: %f\nQtyInStock: %d\n", p.ID, p.Price, p.QtyInStock) // render price and qty
	p.Item.Details()                                                                       // render the item details
}

func (p *Product) DisplayStatus() {
	m := "not in stock"
	if p.IsAvailable() {
		m = "in stock"
	}

	fmt.Printf("Product is %s. Amount left is %d\n", m, p.QtyInStock)
}

func (p *Product) IsAvailable() bool {
	return p.QtyInStock > 0
}
