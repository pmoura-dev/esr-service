package entities

import (
	"net/http"

	"github.com/pmoura-dev/esr-service/internal/handlers/http_handlers"

	"github.com/gin-gonic/gin"
)

func ListEntities(c *gin.Context) {
	entityList, err := http_handlers.EntityService.ListEntities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_handlers.ErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, entityList)
}
