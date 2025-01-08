package entities

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmoura-dev/esr-service/internal/handlers/http_handlers"
	"github.com/pmoura-dev/esr-service/internal/services"
)

func GetEntityByID(c *gin.Context) {
	entityID := c.Param("entity_id")
	if entityID == "" {
		err := errors.New("'entity_id' missing from path")
		c.JSON(http.StatusBadRequest, http_handlers.ErrorMessage(err))
		return
	}

	entity, err := http_handlers.EntityService.GetEntityByID(entityID)
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

	c.JSON(http.StatusOK, entity)
}
