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

func (mlModModel MLModModel) FindAllWithFails() ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("fails > ?", 0).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: No MLModels with fails found`))
	}

	return mlmods, err
}

func (mlModModel MLModModel) FindInOrgNoFails(orgID int) ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("fails == ? AND organization_id = ?", 0, orgID).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels without fails not found in Organization %d`, orgID))
	}

	return mlmods, err
}

func (mlModModel MLModModel) FindAllNoFails() ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("fails == ?", 0).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels without fails not found`))
	}

	return mlmods, err
}

func (mlModModel MLModModel) FindInOrgWithFails(orgID int) ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("fails > ? AND organization_id = ?", 0, orgID).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels not found in Organization %d`, orgID))
	}

	return mlmods, err
}

func (mlModModel MLModModel) FindAllInOrg(orgID int) ([]entities.MLModel, error) {
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("organization_id = ?", orgID).Find(&mlmods)
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
	result := db.Table("mlmodels").Where("Name = ? AND organization_id = ?", mlmodName, orgID).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: MLModels with name %s in Organization %d not found`, mlmodName, orgID))
	}

	return mlmods, err
}

//FindById function looks for MLModels using the ID as an argument
func (mlModModel MLModModel) FindById(mlmodID int) (entities.MLModel, error) {
	var mlmod entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("ID = ?", mlmodID).Find(&mlmod)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf("MLModels not Found with ID %d", mlmodID))
	}

	return mlmod, err
}

//FindById function looks for MLModels in Organization using the ID as an argument
func (mlModModel MLModModel) FindByIdInOrg(mlmodID int, orgID int) (entities.MLModel, error) {
	var mlmod entities.MLModel
	var mlmods []entities.MLModel
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("mlmodels").Where("organization_id = ?", orgID).Find(&mlmods)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf("MLModels not Found with ID %d in Organization %d ", mlmodID, orgID))
	}
	// This is ugly... but still dunno how to restart the ID when a multi condition query is done
	mlmod = mlmods[mlmodID]

	return mlmod, err
}

// CreateMLModel function creates a MLMode in DDBB based in org Name
func (mlModModel MLModModel) CreateMLModel(mlmodName string, orgID uint) (entities.MLModel, error) {
	var mlmod entities.MLModel
	var err error

	// Recover DDBB
	db := database.GetDB(config.CFG.DB.DBType)

	// We don't need to check if the mlmodel exists in the DDBB because the entity allows duplication by name, but we need to check the Organization
	// Create resource
	mlmod = entities.MLModel{
		Name:           mlmodName,
		OrganizationID: orgID,
		Successes:      0,
		Fails:          0,
	}
	result := db.Create(&mlmod)
	if result.RowsAffected == 0 {
		config.MainLogger.Sugar().Panicf("MLModel %s cannot be created in Organization %d", mlmodName, orgID)
		err = errors.New(fmt.Sprintf("MLModel %s cannot be created in Organization %d", mlmodName, orgID))

		return mlmod, err
	}

	return mlmod, err
}
