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
