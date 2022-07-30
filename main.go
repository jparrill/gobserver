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

	database.Identify(cmd.CFG.DB.DBType)

}
