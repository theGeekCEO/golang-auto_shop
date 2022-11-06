package shop

import "fmt"

type Car struct {
	ID    string
	Brand string
	Model string
	Color string
}

func (c *Car) GetID() string {
	return c.ID
}

func (c *Car) GetName() string {
	return fmt.Sprintf("%s-%s (%s)", c.Brand, c.Model, c.Color)
}

func (c *Car) Details() {
	fmt.Printf("Item ID: %s\nCar Brand: %s\nCar Model: %s\nCar Color: %s\n", c.ID, c.Brand, c.Model, c.Color)
}
