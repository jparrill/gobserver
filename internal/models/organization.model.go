package models

import (
	"fmt"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type OrganizationModel struct{}

//FindById function looks for all Organizations in the DDBB
func (organizationModel OrganizationModel) FindAll() ([]entities.Organization, ErrorModel) {
	var orgs []entities.Organization
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	db.Find(&orgs)
	if len(orgs) == 0 {
		err = ErrorModel{
			Msg:  "Organizations not Found in FindAll() function",
			Code: 404,
		}
	}
	return orgs, err
}

//FindById function looks for Organization using the Name as an argument
func (organizationModel OrganizationModel) FindByName(orgName string) (entities.Organization, ErrorModel) {
	var org entities.Organization
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if result.Error != nil {
		err = ErrorModel{
			Msg:  fmt.Sprintf("Organization not Found in FindByName function: %s\n", result.Error),
			Code: 404,
		}
	}
	return org, err
}

//FindById function looks for Organization using the ID as an argument
func (organizationModel OrganizationModel) FindById(orgID uint) (entities.Organization, ErrorModel) {
	var org entities.Organization
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("organizations").Where("id = ?", orgID).Find(&org)
	if result.Error != nil {
		err = ErrorModel{
			Msg:  fmt.Sprintf("Organization not Found in FindById function: %s\n", result.Error),
			Code: 404,
		}
	}
	return org, err
}

// CreateOrg function creates entries in DDBB based on org Name
// It returns the created Organization to be shown
func (organizationModel OrganizationModel) CreateOrg(orgName string) (entities.Organization, ErrorModel) {
	var org entities.Organization
	var err ErrorModel

	// Recover DDBB
	db := database.GetDB(config.CFG.DB.DBType)

	// Check if value exists in DDBB
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name != "" {
		config.MainLogger.Sugar().Errorf("Organization ID already exists and cannot be created: %s\n", org.Name)
		err = ErrorModel{
			Msg:  "Organization cannot be created because already exists",
			Code: 500,
		}

		return org, err

	}

	// Create resource
	org = entities.Organization{
		Name: orgName,
	}
	result := db.Create(&org)
	if result.Error != nil {
		err = ErrorModel{
			Msg:  fmt.Sprintf("Organization cannot be created in CreateOrg function: %s\n", result.Error),
			Code: 500,
		}
	}

	return org, err
}
