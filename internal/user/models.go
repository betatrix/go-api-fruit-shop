package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           uuid.UUID      `json:"id" gorm:"primary_key"`
	UserName     string         `json:"user_name" gorm:"unique"`
	HashPassword float64        `json:"hash_password"`
	Role         string         `json:"role"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
