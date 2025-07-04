package errors

import "errors"

var (
	ErrInvalidFruitID       = errors.New("invalid fruit ID")
	ErrEmptyFruitName       = errors.New("fruit name cannot be empty")
	ErrInvalidFruitPrice    = errors.New("fruit price must be equal or greater than zero")
	ErrInvalidFruitStockQty = errors.New("fruit stock quantity must be equal or greater than zero")
)
