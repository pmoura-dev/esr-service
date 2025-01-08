package entities

import (
	"errors"
	"net/http"

	"github.com/pmoura-dev/esr-service/internal/handlers/http_handlers"
	"github.com/pmoura-dev/esr-service/internal/services"

	"github.com/gin-gonic/gin"
)

func DeleteEntity(c *gin.Context) {
	entityID := c.Param("entity_id")
	if entityID == "" {
		err := errors.New("'entity_id' missing from path")
		c.JSON(http.StatusBadRequest, http_handlers.ErrorMessage(err))
		return
	}

	err := http_handlers.EntityService.DeleteEntity(entityID)
	if err != nil {
		var status int
		switch {
		case errors.Is(err, services.ErrEntityNotFound):
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}

		c.JSON(status, http_handlers.ErrorMessage(err))
		return
	}

	c.Status(http.StatusOK)
}
