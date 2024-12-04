package datastore

import (
	"errors"
)

var (
	ErrConnectionFailed  = errors.New("datastore connection failed")
	ErrInvalidData       = errors.New("data is invalid")
	ErrRecordNotFound    = errors.New("record was not found")
	ErrTableDoesNotExist = errors.New("table does not exist")
	ErrTransactionFailed = errors.New("transaction failed")
)
