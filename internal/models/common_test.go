package models_test

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
)

func setUp() {
	configFile := filepath.Join("testdata", "config", "config.sqlite.yaml")
	config.RecoverConfig(configFile)
	config.InitLogger()
	db := database.Initialize(config.CFG.DB.DBType)
	database.Prepopulate(db, filepath.Join("testdata", "fixtures", "prepopulate_db.json"))
}

func tearDown() {
	dbpath := string(config.CFG.TMPFolder + config.CFG.DB.DBName)
	err := os.Remove(dbpath)
	if err != nil {
		log.Panicf("Error deleting DB file: %s", dbpath)
	}
}
