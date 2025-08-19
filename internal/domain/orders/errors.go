package ordersdomain

import "errors"

var (
	ErrOrderPointNotFound      = errors.New("order point not found")
	ErrOrderPointAlreadyExists = errors.New("order point already exists")
	ErrOrderNotFound           = errors.New("order not found")
	ErrOrderAlreadyExists      = errors.New("order already exists")
)
