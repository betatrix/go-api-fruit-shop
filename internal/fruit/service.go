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

// TODO: ajustar validações
func (s *FruitService) CreateFruits(fruits []FruitRequest) ([]Fruit, error) {
	var name string
	var price float64
	var stockQty int
	var createdFruits []Fruit

	for _, value := range fruits {
		name = *value.Name
		price = *value.Price
		stockQty = *value.StockQty

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

// TODO: ajustar validações
func (s *FruitService) GetFruitbyID(fruitID string) (*Fruit, error) {
	if fruitID == "" {
		return nil, errors.New("invalid data")
	}

	fruit, err := s.repo.GetByID(fruitID)
	if err != nil {
		return nil, err
	}

	return fruit, nil
}

// TODO: ajustar validações
func (s *FruitService) GetAllFruits() (*[]Fruit, error) {
	fruits, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return fruits, nil
}

func (s *FruitService) UpdateFruit(fruitID string, newFruitData FruitRequest) (*Fruit, error) {
	if fruitID == "" {
		return nil, errors.New("invalid fruit ID")
	}

	currentFruit, err := s.repo.GetByID(fruitID)
	if err != nil {
		return nil, err
	}

	if newFruitData.Name != nil {
		if *newFruitData.Name == "" {
			return nil, errors.New("name cannot be empty")
		}
		currentFruit.Name = *newFruitData.Name
	}
	if newFruitData.Price != nil {
		if *newFruitData.Price <= 0 {
			return nil, errors.New("price cannot be equal to or less than zero")
		}
		currentFruit.Price = *newFruitData.Price
	}
	if newFruitData.StockQty != nil {
		if *newFruitData.StockQty <= 0 {
			return nil, errors.New("stock quantity cannot be equal to or less than zero")
		}
		currentFruit.StockQty = *newFruitData.StockQty
	}

	err = s.repo.Update(currentFruit)
	if err != nil {
		return nil, err
	}

	return currentFruit, nil
}
