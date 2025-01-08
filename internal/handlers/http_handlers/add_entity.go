package http_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

func AddEntity(c *gin.Context) {
	var entity models.Entity

	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, errorMessage(ErrBadRequest))
		return
	}

	if err := entityService.AddEntity(entity); err != nil {
		c.JSON(http.StatusInternalServerError, errorMessage(err))
		return
	}

	c.JSON(http.StatusCreated, resultMessage("Success", nil))
}
