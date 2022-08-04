package models

import (
	"fmt"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type MLModModel struct{}

func (mlModModel MLModModel) FindAll() ([]entities.MLModel, ErrorModel) {
	var mlmods []entities.MLModel
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	db.Find(&mlmods)
	if len(mlmods) == 0 {
		err = ErrorModel{
			Msg:  "MLModels not Found in FindAll() function",
			Code: 404,
		}
	}
	return mlmods, err
}

//FindById function looks for MLModels using the Name as an argument
func (mlModModel MLModModel) FindByName(mlmodName string) ([]entities.MLModel, ErrorModel) {
	var mlmods []entities.MLModel
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	db.Table("mlmodels").Where("Name = ?", mlmodName).Find(&mlmods)
	if len(mlmods) == 0 {
		err = ErrorModel{
			Msg:  "MLModels not Found in FindByName() function",
			Code: 404,
		}
	}
	return mlmods, err
}

//FindById function looks for Organization using the ID as an argument
func (mlModModel MLModModel) FindById(mlmodID uint) (entities.MLModel, ErrorModel) {
	var mlmod entities.MLModel
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("id = ?", mlmodID).Find(&mlmod)
	if result.Error != nil {
		err = ErrorModel{
			Msg:  fmt.Sprintf("MLModels not Found in FindById function: %s\n", result.Error),
			Code: 404,
		}
	}
	return mlmod, err
}

// CreateOrg function creates entries in DDBB based on org Name
func (mlModModel MLModModel) CreateMLModel(mlmodName string, orgName string) (entities.MLModel, ErrorModel) {
	var mlmod entities.MLModel
	var org entities.Organization
	var err ErrorModel

	// Recover DDBB
	db := database.GetDB(config.CFG.DB.DBType)

	// We don't need to check if the mlmodel exists in the DDBB because the entity allows duplication by name, but we need to check the Organization
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name == "" {
		config.MainLogger.Sugar().Errorf("Organization does not exists and cannot be associated with the MLModel: %s\n", orgName)
		err = ErrorModel{
			Msg:  fmt.Sprintf("Organization does not exists and cannot be associated with the MLModel: %s\n", orgName),
			Code: 500,
		}

		return mlmod, err

	}

	// Create resource
	mlmod = entities.MLModel{
		Name:           mlmodName,
		OrganizationID: org.ID,
		Successes:      0,
		Fails:          0,
	}
	result := db.Create(&mlmod)
	if result.Error != nil {
		err = ErrorModel{
			Msg:  fmt.Sprintf("MLModel cannot be created in CreateMLModel function: %s\n", result.Error),
			Code: 500,
		}
		return mlmod, err
	}

	return mlmod, err
}
