package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
