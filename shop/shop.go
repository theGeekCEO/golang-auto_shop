package shop

func NewStore() *Store {
	return &Store{
		ID:    GenerateID(),
		Stock: make(map[string]Product),
	}
}

func NewCar(brand, model, color string) *Car {
	return &Car{
		ID:    GenerateID(),
		Brand: brand,
		Model: model,
		Color: color,
	}
}
