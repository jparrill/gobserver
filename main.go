package main

import (
	"context"
	"fmt"

	"github.com/jparrill/gobserver/internal/cmd"
)

func main() {

	// Initialize Context
	ctx := context.Background()

	// Initialize the logger
	logger := cmd.InitLogger(ctx)
	logger.Info("Loading Configuration")

	// Recover the configuration
	cfg := cmd.RecoverConfig(ctx)
	fmt.Printf("%T, %v", cfg, cfg.DB.DBName)

}
