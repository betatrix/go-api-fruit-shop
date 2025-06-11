package fruit

import (
	"github.com/betatrix/go-api-fruit-shop/internal/errors"
)

type FruitService struct {
	repo *FruitRepository
}

func NewFruitService(repo *FruitRepository) *FruitService {
	return &FruitService{repo: repo}
}

func (s *FruitService) CreateFruits(fruits FruitsDTO) ([]Fruit, error) {
	var createdFruits []Fruit

	for _, value := range fruits.Fruits {
		if value.Name == nil || *value.Name == "" {
			return nil, errors.ErrEmptyFruitName
		}
		if value.Price == nil || *value.Price <= 0 {
			return nil, errors.ErrInvalidFruitPrice
		}
		if value.StockQty == nil || *value.StockQty <= 0 {
			return nil, errors.ErrInvalidFruitStockQty
		}

		fruit := NewFruitModel(*value.Name, *value.Price, *value.StockQty)

		err := s.repo.Create(&fruit)
		if err != nil {
			return nil, err
		}

		createdFruits = append(createdFruits, fruit)
	}

	return createdFruits, nil
}

func (s *FruitService) GetFruitbyID(fruitID string) (*Fruit, error) {
	if fruitID == "" {
		return nil, errors.ErrInvalidFruitID
	}

	fruit, err := s.repo.GetByID(fruitID)
	if err != nil {
		return nil, err
	}

	return fruit, nil
}

func (s *FruitService) GetAllFruits() (*[]Fruit, error) {
	fruits, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return fruits, nil
}

func (s *FruitService) UpdateFruit(fruitID string, newFruitData FruitDTO) (*Fruit, error) {
	if fruitID == "" {
		return nil, errors.ErrInvalidFruitID
	}

	currentFruit, err := s.repo.GetByID(fruitID)
	if err != nil {
		return nil, err
	}

	if newFruitData.Name != nil {
		if *newFruitData.Name == "" {
			return nil, errors.ErrEmptyFruitName
		}
		currentFruit.Name = *newFruitData.Name
	}
	if newFruitData.Price != nil {
		if *newFruitData.Price <= 0 {
			return nil, errors.ErrInvalidFruitPrice
		}
		currentFruit.Price = *newFruitData.Price
	}
	if newFruitData.StockQty != nil {
		if *newFruitData.StockQty <= 0 {
			return nil, errors.ErrInvalidFruitStockQty
		}
		currentFruit.StockQty = *newFruitData.StockQty
	}

	err = s.repo.Update(currentFruit)
	if err != nil {
		return nil, err
	}

	return currentFruit, nil
}

// TODO: ajustar validações
func (s *FruitService) DeleteFruit(fruitID string) (string, error) {
	err := s.repo.Delete(fruitID)
	if err != nil {
		return "", err
	}

	successMsg := "Fruit ID:" + fruitID + " - Deleted with success!"

	return successMsg, nil
}
