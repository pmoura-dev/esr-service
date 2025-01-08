package services

import (
	"errors"
)

var (
	ErrEntityNotFound      = errors.New("entity not found")
	ErrEntityAlreadyExists = errors.New("entity already exists")
	ErrInternalError       = errors.New("internal error")
)
