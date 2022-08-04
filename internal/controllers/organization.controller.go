package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/entities"
	"github.com/jparrill/gobserver/internal/models"
)

type OrganizationController struct{}

var om models.OrganizationModel

// Retrieve function recover the org details based on Name or ID.
func (oc OrganizationController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
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

	} else if c.Param("name") != "" {
		orgName := string(c.Param("name"))
		org, err := om.FindByName(orgName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Error to retrieving Organization", "error": err})
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

// Create function receives a slice of JSON, to be more concrete slice of Organization
// which only implies the name. If the organization exists, will not be recreated.
func (oc OrganizationController) Create(c *gin.Context) {
	var orgs, orgsCreated []entities.Organization
	var orgModel models.OrganizationModel

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		config.MainLogger.Sugar().Panicf("Error recovering data from Request: %v", err)
	}
	err = json.Unmarshal(jsonData, &orgs)
	if err != nil {
		config.MainLogger.Sugar().Panicf("Error unmarshalling JSON payload: %v", err)
	}

	for _, org := range orgs {
		if !(orgModel.OrgExists(org.Name)) {
			o, err := orgModel.CreateOrg(org.Name)
			if err != nil {
				config.MainLogger.Sugar().Panicf(fmt.Sprintf(`Organization "%s" cannot be created`, org.Name))
			}
			config.MainLogger.Sugar().Info(fmt.Sprintf(`Organization "%s" created`, org.Name))
			orgsCreated = append(orgsCreated, o)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organizations created", "Organizations": orgsCreated})
	return
}
