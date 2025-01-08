package http_handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/pmoura-dev/esr-service/internal/services"
	"github.com/pmoura-dev/esr-service/internal/validation"
)

var (
	EntityService services.EntityService
)

var (
	ErrInvalidJSONBody = errors.New("invalid JSON body")
	ErrBadRequest      = errors.New("bad request")
	ErrInternalError   = errors.New("internal error")
)

func ResultMessage(message string, params map[string]any) gin.H {
	if len(params) == 0 {
		return gin.H{"message": message}
	}

	return gin.H{"message": message, "params": params}
}

func ErrorMessage(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ValidationErrorMessage(errorList validation.ErrorList) gin.H {
	return gin.H{
		"error":   "Validation Error",
		"details": errorList,
	}
}
