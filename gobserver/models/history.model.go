package models

import (
	"errors"
	"fmt"

	"github.com/jparrill/gobserver/config"
	"github.com/jparrill/gobserver/database"
	"github.com/jparrill/gobserver/entities"
)

type HistoryModel struct{}

func (historyModel HistoryModel) FindAll() ([]entities.History, error) {
	var hists []entities.History
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	result := db.Find(&hists)
	if result.RowsAffected == 0 {
		err = errors.New("History entries not Found in FindAll() function")
	}
	return hists, err
}

func (historyModel HistoryModel) FindMlModelHistory(MLModelID int, orgName string) ([]entities.History, error) {
	var hists []entities.History
	var org entities.Organization
	var err error

	db := database.GetDB(config.CFG.DB.DBType)
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name == "" {
		config.MainLogger.Sugar().Panicf("Organization does not exists: %s\n", orgName)
		err = errors.New(fmt.Sprintf("Organization does not exists: %s\n", orgName))

		return nil, err
	}

	result := db.Table("history").Where("OrganizationID = ? AND MLModelID = ?", org.ID, MLModelID)
	if result.RowsAffected == 0 {
		err = errors.New(fmt.Sprintf("History over MLModel %d in Organization %s not found: %s\n", MLModelID, orgName, result.Error))
		return hists, err
	}

	db.Find(&hists)
	return hists, err
}
