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

func (r *UserRepository) GetByID(userID string) (*User, error) {
	var user User

	result := r.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetAll() (*[]User, error) {
	var users []User

	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}
