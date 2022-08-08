package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jparrill/gobserver/config"
	"github.com/jparrill/gobserver/entities"
	"github.com/jparrill/gobserver/models"
)

type MLModelController struct{}

var mlmm models.MLModModel

func (mlmc MLModelController) RetrieveId(c *gin.Context) {
	orgID, e := strconv.ParseUint(c.Param("orgId"), 10, 64)
	if e != nil {
		config.MainLogger.Sugar().Panicf("Error parsing Organization ID: %v", e)
	}

	mlmodelID, e := strconv.ParseUint(c.Param("id"), 10, 64)
	if e != nil {
		config.MainLogger.Sugar().Panicf("Error parsing MLModel ID: %v", e)
	}

	mlmod, err := mlmm.FindByIdInOrg(int(mlmodelID), int(orgID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Error retrieving MLModel %d in Organization %d`, mlmodelID, orgID), "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "MLModel found!", "MLModel": mlmod})
	return
}

func (mlmc MLModelController) RetrieveModNoFails(c *gin.Context) {
	mlmod, err := mlmm.FindAllNoFails()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Error retrieving MLModels without fails`), "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All MLModel without fails", "MLModel": mlmod})
	return
}

func (mlmc MLModelController) RetrieveModInOrgNoFails(c *gin.Context) {
	orgID, e := strconv.ParseUint(c.Param("orgId"), 10, 64)
	if e != nil {
		config.MainLogger.Sugar().Panicf("Error parsing Organization ID: %v", e)
	}

	mlmod, err := mlmm.FindInOrgNoFails(int(orgID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`Error retrieving MLModels in Organization %d`, orgID), "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All MLModel without fails", "MLModel": mlmod})
	return
}

// Create function receives a JSON with the details about the MLmodel Name and organization_id which belongs to
func (mlmc MLModelController) Create(c *gin.Context) {
	var orgModel models.OrganizationModel
	var mlModModel models.MLModModel
	var mlModel entities.MLModel

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		config.MainLogger.Sugar().Panicf("Error recovering data from Request: %v", err)
	}
	err = json.Unmarshal(jsonData, &mlModel)
	if err != nil {
		config.MainLogger.Sugar().Panicf("Error unmarshalling JSON payload: %v", err)
	}

	if !(orgModel.OrgExists(mlModel.Name)) {
		m, err := mlModModel.CreateMLModel(mlModel.Name, uint(mlModel.OrganizationID))
		if err != nil {
			config.MainLogger.Sugar().Panicf(fmt.Sprintf(`MLModel "%s" cannot be created in %d Organization`, mlModel.Name, mlModel.OrganizationID))
		}
		config.MainLogger.Sugar().Info(fmt.Sprintf(`MLModel "%s" created into %d OrganizationID`, m.Name, m.OrganizationID))
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf(`"MLModel %s created in Organization %d"`, mlModel.Name, mlModel.OrganizationID), "MLModel": mlModel})

	return
}

// CreateBulk function receives a JSON with the details about the MLmodel Name and organization_id which belongs to
func (mlmc MLModelController) CreateBulk(c *gin.Context) {
	var orgModel models.OrganizationModel
	var mlModModel models.MLModModel
	var mlModel, mlModelsCreated []entities.MLModel

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		config.MainLogger.Sugar().Panicf("Error recovering data from Request: %v", err)
	}
	err = json.Unmarshal(jsonData, &mlModel)
	if err != nil {
		config.MainLogger.Sugar().Panicf("Error unmarshalling JSON payload: %v", err)
	}

	for i, k := range mlModel {
		if !(orgModel.OrgExists(k.Name)) {
			m, err := mlModModel.CreateMLModel(k.Name, uint(k.OrganizationID))
			if err != nil {
				config.MainLogger.Sugar().Panicf(fmt.Sprintf(`Index %d MLModel "%s" cannot be created in %d Organization`, i, k.Name, k.OrganizationID))
			}
			config.MainLogger.Sugar().Info(fmt.Sprintf(`MLModel "%s" created into %d OrganizationID`, m.Name, m.OrganizationID))
			mlModelsCreated = append(mlModelsCreated, m)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf(`"MLModels created %v"`, mlModelsCreated), "MLModels": mlModelsCreated})
	return
}
