package models

import (
	"errors"
	"fmt"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type OrganizationModel struct{}

//FindById function looks for all Organizations in the DDBB
func (organizationModel OrganizationModel) FindAll() ([]entities.Organization, error) {
	var orgs []entities.Organization
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Find(&orgs)
	if result.RowsAffected == 0 {
		err = errors.New(`Error: Organizations not found`)
	}
	return orgs, err
}

func (organizationModel OrganizationModel) OrgExists(orgName string) bool {
	var org entities.Organization

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

//FindById function looks for Organization using the Name as an argument
func (organizationModel OrganizationModel) FindByName(orgName string) (entities.Organization, error) {
	var org entities.Organization
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: Organization with name "%s" not found`, orgName))
	}
	return org, err
}

//FindById function looks for Organization using the ID as an argument
func (organizationModel OrganizationModel) FindById(orgID uint) (entities.Organization, error) {
	var org entities.Organization
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Table("organizations").Where("id = ?", orgID).Find(&org)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf(`Error: Organization with id "%d" not found`, orgID))
	}
	return org, err
}

// CreateOrg function creates entries in DDBB based on org Name
// It returns the created Organization to be shown
func (organizationModel OrganizationModel) CreateOrg(orgName string) (entities.Organization, error) {
	var org entities.Organization
	var err error

	// Recover DDBB
	db := database.GetDB(config.CFG.DB.DBType)

	// Check if value exists in DDBB
	result := db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name != "" {
		config.MainLogger.Sugar().Panicf("Organization ID already exists and cannot be created: %s\n", org.Name)
		err = errors.New(fmt.Sprintf("Organization ID already exists and cannot be created: %s\n", org.Name))
		return org, err
	}

	// Create resource
	org = entities.Organization{
		Name: orgName,
	}
	result = db.Create(&org)
	if result.RowsAffected == 0 {
		config.MainLogger.Sugar().Panicf("Organization %s cannot be created", orgName)
		err = errors.New(fmt.Sprintf("Organization %s cannot be created", orgName))
		return org, err
	}

	return org, err
}
