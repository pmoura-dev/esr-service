package http_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListEntities(c *gin.Context) {
	entityList, err := entityService.ListEntities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorMessage(ErrInternalError))
	}

	c.JSON(http.StatusOK, resultMessage("Success", map[string]any{
		"entityList": entityList,
	}))
}
