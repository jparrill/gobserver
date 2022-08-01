package models

import (
	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type OrganizationModel struct{}

func (organizationModel OrganizationModel) FindAll() []entities.Organization {
	var orgs []entities.Organization

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Find(&orgs)
	return orgs
}

func (organizationModel OrganizationModel) FindByName(orgName string) entities.Organization {
	var org entities.Organization

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	return org
}

func (organizationModel OrganizationModel) FindById(orgID uint) entities.Organization {
	var org entities.Organization

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Table("organizations").Where("id = ?", orgID).Find(&org)
	return org
}

// CreateOrg function creates entries in DDBB based on org Name
func (organizationModel OrganizationModel) CreateOrg(orgName string) {
	var org entities.Organization

	// Recover DDBB
	db := database.GetDB(cmd.CFG.DB.DBType)

	// Check if value exists in DDBB
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name != "" {
		cmd.MainLogger.Sugar().Panicf("Organization ID already exists: %s\n", org.Name)
	}

	// Create resource
	org = entities.Organization{
		Name: orgName,
	}
	db.Create(&org)
}
