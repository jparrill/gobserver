package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/models"
)

type MLModelController struct{}

var mlmodmodel models.MLModModel

func (mlmc MLModelController) Retrieve(c *gin.Context) {
	orgID, e := strconv.ParseUint(c.Param("id"), 10, 64)
	if e != nil {
		config.MainLogger.Sugar().Panicf("Error parsing Organization ID: %v", e)
	}

	org, err := om.FindById(uint(orgID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Error retrieving Organization", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization found!", "Organization": org})
	return

}
