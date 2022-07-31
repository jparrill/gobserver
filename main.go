package main

import (
	"context"

	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/database"
)

var ctx context.Context

func main() {

	// Initialize Context
	ctx = context.Background()

	// Recover the configuration
	cmd.RecoverConfig()

	// Initialize the logger
	cmd.InitLogger()
	cmd.MainLogger.Info("Initializing DDBB")

	// Recovering DB handler from initialization
	db := database.Initialize(cmd.CFG.DB.DBType)

	database.Prepopulate(db)
	//database.Query(db)

}
