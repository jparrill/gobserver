package models

import (
	"fmt"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/entities"
)

type HistoryModel struct{}

func (historyModel HistoryModel) FindAll() ([]entities.History, ErrorModel) {
	var hists []entities.History
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	db.Find(&hists)
	if len(hists) == 0 {
		err = ErrorModel{
			Msg:  "History entries not Found in FindAll() function",
			Code: 404,
		}
	}
	return hists, err
}

func (historyModel HistoryModel) FindMlModelHistory(MLModelID int, orgName string) ([]entities.History, ErrorModel) {
	var hists []entities.History
	var org entities.Organization
	var err ErrorModel

	db := database.GetDB(config.CFG.DB.DBType)
	db.Table("organizations").Where("Name = ?", orgName).Find(&org)
	if org.Name == "" {
		config.MainLogger.Sugar().Panicf("Organization does not exists: %s\n", orgName)
		err = ErrorModel{
			Msg:  fmt.Sprintf("Organization does not exists: %s\n", orgName),
			Code: 404,
		}
		return nil, err

	}

	result := db.Table("history").Where("OrganizationID = ? AND MLModelID = ?", org.ID, MLModelID)
	if result.Error != nil {
		err = ErrorModel{
			Msg:  fmt.Sprintf("History over MLModel %d in Organization %s not found: %s\n", MLModelID, orgName, result.Error),
			Code: 404,
		}
		return hists, err
	}

	db.Find(&hists)
	return hists, err
}
