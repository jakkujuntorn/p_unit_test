package service

import "errors"

var (
	ErrZeroAmount = errors.New("Purchase amount is zero")
	ErrRepository = errors.New("Repository Error")
)