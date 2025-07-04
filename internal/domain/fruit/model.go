package fruit

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Fruit struct {
	ID        string         `uri:"id" json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	StockQty  int            `json:"stock_qty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type FruitDTO struct {
	Name     *string  `json:"name"`
	Price    *float64 `json:"price"`
	StockQty *int     `json:"stock_qty"`
}

type FruitsDTO struct {
	Fruits []FruitDTO `json:"fruits"`
}

func NewFruitModel(name string, price float64, stockQty int) Fruit {
	return Fruit{
		ID:        uuid.New().String(),
		Name:      name,
		Price:     price,
		StockQty:  stockQty,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
