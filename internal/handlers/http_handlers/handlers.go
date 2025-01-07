package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pmoura-dev/esr-service/internal/services"
)

var (
	entityService services.EntityService
)

func Setup(es services.EntityService) {
	entityService = es
}

func resultMessage(message string, params map[string]any) gin.H {
	if len(params) == 0 {
		return gin.H{"message": message}
	}

	return gin.H{"message": message, "params": params}
}

func errorMessage(err error) gin.H {
	return gin.H{"message": err.Error()}
}
