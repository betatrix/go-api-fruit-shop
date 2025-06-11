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

// Creates a new fruit
func (r *FruitRepository) Create(fruit *Fruit) error {
	result := r.db.Create(&fruit)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Searches and returns the fruit by id
func (r *FruitRepository) GetByID(fruitID string) (*Fruit, error) {
	var fruit Fruit

	result := r.db.First(&fruit, "id = ?", fruitID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &fruit, nil
}

// List all fruits
func (r *FruitRepository) GetAll() (*[]Fruit, error) {
	var fruits []Fruit

	result := r.db.Find(&fruits)
	if result.Error != nil {
		return nil, result.Error
	}

	return &fruits, nil
}

// Updates fields passed by the user
func (r *FruitRepository) Update(fruit *Fruit) error {
	result := r.db.Model(&fruit).Updates(fruit)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Performs a soft delete (DeletedAt field included in a Fruit struct)
func (r *FruitRepository) Delete(fruitID string) error {
	var fruit Fruit

	result := r.db.Where("id = ?", fruitID).Delete(&fruit)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
