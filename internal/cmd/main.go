package main

import (
	"context"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/server"
)

var ctx context.Context

func main() {

	// Initialize Context
	ctx = context.Background()

	// Recover the configuration
	config.RecoverConfig("config.yaml")
	// Initialize the logger
	config.InitLogger()
	config.MainLogger.Info("Initializing DDBB")

	// Recovering DB handler from initialization
	db := database.Initialize(config.CFG.DB.DBType)
	database.Prepopulate(db, "fixtures/prepopulate_db.json")

	// Run Server
	server.Init()
}
