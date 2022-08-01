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
