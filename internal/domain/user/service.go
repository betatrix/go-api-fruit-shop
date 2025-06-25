package user

import (
	"unicode/utf8"

	"github.com/betatrix/go-api-fruit-shop/internal/enums"
	"github.com/betatrix/go-api-fruit-shop/internal/errors"
	"github.com/betatrix/go-api-fruit-shop/internal/utils"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(userReq UserDTO) (*User, error) {
	if userReq.Username == nil || *userReq.Username == "" {
		return nil, errors.ErrEmptyUsername
	}
	if utf8.RuneCountInString(*userReq.Username) < 3 {
		return nil, errors.ErrInvalidUsername
	}
	if userReq.Password == nil {
		return nil, errors.ErrEmptyPassword
	}
	if utf8.RuneCountInString(*userReq.Password) < 5 {
		return nil, errors.ErrInvalidPassword
	}

	role := enums.ParseRole(*userReq.Role)
	if role == "" {
		return nil, errors.ErrInvalidUserRole
	}

	hashPassword, err := utils.HashPassword(*userReq.Password)
	if err != nil {
		return nil, err
	}

	user := NewUserModel(*userReq.Username, hashPassword, role)

	err = s.repo.Create(&user)
	if err != nil {
		if errors.IsDuplicateEntry(err) {
			return nil, errors.ErrUsernameAlreadyExists
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetUserbyID(userID string) (*User, error) {
	if userID == "" {
		return nil, errors.ErrInvalidUserID
	}

	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() (*[]User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
