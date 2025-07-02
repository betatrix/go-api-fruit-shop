package errors

import "errors"

var (
	ErrInvalidUserID          = errors.New("invalid user ID")
	ErrEmptyUsername          = errors.New("username cannot be empty")
	ErrUsernameAlreadyExists  = errors.New("username already exists, please choose another")
	ErrInvalidUsername        = errors.New("username must be greater than three digits")
	ErrEmptyPassword          = errors.New("user password cannot be empty")
	ErrInvalidPassword        = errors.New("user password must be greater than five digits")
	ErrInvalidUserRole        = errors.New("user role must be 'admin' or 'user'")
	ErrInvalidUserCredentials = errors.New("invalid username or password, please try again")
)
