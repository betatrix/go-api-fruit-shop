package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Creates a new user
func (r *UserRepository) Create(userReq *User) error {
	result := r.db.Create(&userReq)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Searches and returns the user by id
func (r *UserRepository) GetByID(userID string) (*User, error) {
	var user User

	result := r.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// List all users
func (r *UserRepository) GetAll() (*[]User, error) {
	var users []User

	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

// List all users
func (r *UserRepository) FindUser(userReq *UserLoginDTO) (*User, error) {
	var user User

	result := r.db.First(&user, "username = ?", userReq.Username)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
