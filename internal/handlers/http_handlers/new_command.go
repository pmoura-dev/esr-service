package http_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewCommand(c *gin.Context) {
	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, errorMessage(ErrRequiredEntityID))
	}

	var desiredState map[string]any
	if err := c.ShouldBindJSON(&desiredState); err != nil {
		c.JSON(http.StatusBadRequest, errorMessage(ErrBadRequest))
	}

	commandID, err := entityService.ProcessCommand(entityID, desiredState)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorMessage(err))
	}

	c.JSON(http.StatusAccepted, resultMessage("Command was accepted.", map[string]any{
		"command_id": commandID,
	}))
}
