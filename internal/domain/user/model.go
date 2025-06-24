package user

import (
	"time"

	"github.com/betatrix/go-api-fruit-shop/internal/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string         `json:"id" gorm:"primary_key"`
	Username     string         `json:"username" gorm:"unique"`
	HashPassword string         `json:"hash_password"`
	Role         enums.Role     `json:"role" gorm:"type:enum('admin','user');default:'user'"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type UserDTO struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Role     *string `json:"role"`
}

func NewUserModel(username string, password string, role enums.Role) User {
	return User{
		ID:           uuid.New().String(),
		Username:     username,
		HashPassword: password,
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
