package fruit

import (
	"gorm.io/gorm"
)

type FruitRepository struct {
	db *gorm.DB
}

func NewFruitRepository(db *gorm.DB) *FruitRepository {
	return &FruitRepository{db: db}
}

func (r *FruitRepository) Create(fruit *Fruit) error {
	result := r.db.Create(&fruit)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
