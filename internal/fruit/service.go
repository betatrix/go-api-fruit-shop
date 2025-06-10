package fruit

import (
	"errors"
)

type FruitService struct {
	repo *FruitRepository
}

func NewFruitService(repo *FruitRepository) *FruitService {
	return &FruitService{repo: repo}
}

func (s *FruitService) CreateFruits(fruits []FruitRequest) ([]Fruit, error) {
	var name string
	var price float64
	var stockQty int
	var createdFruits []Fruit

	for _, value := range fruits {
		name = value.Name
		price = value.Price
		stockQty = value.StockQty

		if name == "" || price < 0 || stockQty < 0 {
			return nil, errors.New("invalid data")
		}

		fruit := NewFruitModel(name, price, stockQty)

		err := s.repo.Create(&fruit)
		if err != nil {
			return nil, err
		}

		createdFruits = append(createdFruits, fruit)
	}

	return createdFruits, nil
}
