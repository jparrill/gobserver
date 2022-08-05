package models

import (
	"errors"
	"fmt"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type MLModModel struct{}

func (mlModModel MLModModel) FindAll() ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(`Error: MLModels not found`)
	}
	return mlmods, err
}

func (mlModModel MLModModel) FindAllInOrg(orgID int) ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("OrganizationID = ?", orgID).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels not found in Organization %d`, orgID))
	}

	return mlmods, err
}

//FindById function looks for MLModels using the Name as an argument
func (mlModModel MLModModel) FindByName(mlmodName string) ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("Name = ?", mlmodName).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels with name %s not found`, mlmodName))
	}
	return mlmods, err
}

//FindById function looks for MLModels in a Organization using the Name as an argument
func (mlModModel MLModModel) FindByNameInOrg(mlmodName string, orgID int) ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("Name = ? AND OrganizationID = ?", mlmodName, orgID).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels with name %s in Organization %d not found`, mlmodName, orgID))
	}

	return mlmods, err
}

//FindById function looks for MLModels using the ID as an argument
func (mlModModel MLModModel) FindById(mlmodID uint) (entities.MLModel, error) {
	var mlmod entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("id = ?", mlmodID).Find(&mlmod)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf("MLModels not Found with ID %d", mlmodID))
	}

	return mlmod, err
}

//FindById function looks for MLModels in Organization using the ID as an argument
func (mlModModel MLModModel) FindByIdInOrg(mlmodID uint, orgID int) (entities.MLModel, error) {
	var mlmod entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("id = ? AND OrganizationID = ?", mlmodID, orgID).Find(&mlmod)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf("MLModels not Found with ID %d in Organization %d ", mlmodID, orgID))
	}

	return mlmod, err
}

// CreateMLModel function creates a MLMode in DDBB based in org Name
func (mlModModel MLModModel) CreateMLModel(mlmodName string, orgName string) (entities.MLModel, error) {
	var mlmod entities.MLModel
	var org entities.Organization
	var err error

	// Recover DDBB
	db := database.GetDB(config.CFG.DB.DBType)

	// We don't need to check if the mlmodel exists in the DDBB because the entity allows duplication by name, but we need to check the Organization
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name == "" {
		config.MainLogger.Sugar().Panicf("Organization does not exists and cannot be associated with the MLModel: %s\n", orgName)
		err = errors.New(fmt.Sprintf("Organization does not exists and cannot be associated with the MLModel: %s\n", orgName))

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
	if result.RowsAffected == 0 {
		config.MainLogger.Sugar().Panicf("MLModel %s cannot be created in Organization %s", mlmodName, orgName)
		err = errors.New(fmt.Sprintf("MLModel %s cannot be created in Organization %s", mlmodName, orgName))

		return mlmod, err
	}

	return mlmod, err
}
