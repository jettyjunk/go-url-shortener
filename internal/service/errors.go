package service

import "errors"

var (
	ErrInvalidURL     = errors.New("invaled url")
	ErrCodeGeneration = errors.New("failed to generate unique code")
)
