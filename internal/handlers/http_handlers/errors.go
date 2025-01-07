package http_handlers

import (
	"errors"
)

var (
	ErrRequiredEntityID = errors.New("entity_id is required")
	ErrBadRequest       = errors.New("bad request")
	ErrInternalError    = errors.New("internal error")
)
