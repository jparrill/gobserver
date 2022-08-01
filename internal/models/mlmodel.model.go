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
	return mlmodels
}
