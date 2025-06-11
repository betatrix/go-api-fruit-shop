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

func (r *FruitRepository) GetByID(fruitID string) (*Fruit, error) {
	var fruit Fruit

	result := r.db.First(&fruit, "id = ?", fruitID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &fruit, nil
}

func (r *FruitRepository) GetAll() (*[]Fruit, error) {
	var fruits []Fruit

	result := r.db.Find(&fruits)
	if result.Error != nil {
		return nil, result.Error
	}

	return &fruits, nil
}

func (r *FruitRepository) Update(fruit *Fruit) error {
	result := r.db.Model(&fruit).Updates(fruit)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
