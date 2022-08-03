package models

import (
	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type MLModModel struct{}

func (mlModModel MLModModel) FindAll() []entities.MLModel {
	var mlmodels []entities.MLModel

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Find(&mlmodels)
	db.Model(&entities.MLModel{}).Preload("Organization").Find(&mlmodels)
	return mlmodels
}

//FindById function looks for MLModels using the Name as an argument
func (mlModModel MLModModel) FindByName(mlmodName string) []entities.MLModel {
	var mlmods []entities.MLModel

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Table("mlmodels").Where("Name = ?", mlmodName).Find(&mlmods)
	return mlmods
}

//FindById function looks for Organization using the ID as an argument
func (mlModModel MLModModel) FindById(mlmodID uint) entities.MLModel {
	var mlmod entities.MLModel

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Table("mlmodels").Where("id = ?", mlmodID).Find(&mlmod)
	return mlmod
}

// CreateOrg function creates entries in DDBB based on org Name
func (mlModModel MLModModel) CreateMLModel(mlmodName string, orgName string) entities.MLModel {
	var mlmod entities.MLModel
	var org entities.Organization

	// Recover DDBB
	db := database.GetDB(cmd.CFG.DB.DBType)

	// We don't need to check if the mlmodel exists in the DDBB because the entity allows duplication by name, but we need to check the Organization
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name == "" {
		cmd.MainLogger.Sugar().Panicf("Organization does not exists: %s\n", orgName)
	}

	// Create resource
	mlmod = entities.MLModel{
		Name:           mlmodName,
		OrganizationID: org.ID,
		Successes:      0,
		Fails:          0,
	}
	db.Create(&mlmod)

	return mlmod
}
