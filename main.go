package main

import (
	"context"

	"github.com/jparrill/gobserver/internal/cmd"
)

var ctx context.Context

func main() {

	// Initialize Context
	ctx = context.Background()

	// Recover the configuration
	cmd.RecoverConfig()

	// Initialize the logger
	logger := cmd.InitLogger()
	logger.Info("Loading Configuration")

}
