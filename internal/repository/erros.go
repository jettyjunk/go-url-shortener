package repository

import "errors"

var (
	ErrCodeExists = errors.New("code already exists")
	ErrNotFound   = errors.New("link not found")
)
