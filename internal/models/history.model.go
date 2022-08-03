package models

import (
	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type HistoryModel struct{}

func (historyModel HistoryModel) FindAll() []entities.History {
	var hist []entities.History

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Find(&hist)
	return hist
}

func (historyModel HistoryModel) FindMlModelHistory(MLModelID int, orgName string) []entities.History {
	var hist []entities.History
	var org entities.Organization

	db := database.GetDB(cmd.CFG.DB.DBType)
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name == "" {
		cmd.MainLogger.Sugar().Panicf("Organization does not exists: %s\n", orgName)
	}

	db.Table("history").Where("OrganizationID = ? AND MLModelID = ?", org.ID, MLModelID)

	db.Find(&hist)
	return hist
}
