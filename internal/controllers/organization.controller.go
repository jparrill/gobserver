package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/models"
)

type OrganizationController struct{}

var om models.OrganizationModel

func (oc OrganizationController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		orgID, e := strconv.ParseUint(c.Param("id"), 10, 64)
		if e != nil {
			config.MainLogger.Sugar().Errorf("Error parsing Organization ID: %v", e)
		}

		org, err := om.FindById(uint(orgID))
		if err.Code == 404 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieving Organization", "error": err.Msg})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Organization found!", "Organization": org})
		return

	} else if c.Param("name") != "" {
		org, err := om.FindByName(c.Param("name"))
		if err.Code == 404 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieving Organization", "error": err.Msg})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Organization found!", "Organization": org})
		return

	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
