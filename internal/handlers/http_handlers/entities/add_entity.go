package entities

import (
	"errors"
	"net/http"

	"github.com/pmoura-dev/esr-service/internal/handlers/http_handlers"
	"github.com/pmoura-dev/esr-service/internal/services"
	"github.com/pmoura-dev/esr-service/internal/types"

	"github.com/gin-gonic/gin"
)

func AddEntity(c *gin.Context) {
	var entity types.Entity

	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, http_handlers.ErrorMessage(http_handlers.ErrInvalidJSONBody))
		return
	}

	if errorList := entity.Validate(); len(errorList) > 0 {
		c.JSON(http.StatusBadRequest, http_handlers.ValidationErrorMessage(errorList))
		return
	}

	if err := http_handlers.EntityService.AddEntity(entity); err != nil {
		var status int
		switch {
		case errors.Is(err, services.ErrEntityAlreadyExists):
			status = http.StatusConflict
		default:
			status = http.StatusInternalServerError
		}

		c.JSON(status, http_handlers.ErrorMessage(err))
		return
	}

	c.Status(http.StatusCreated)
}
