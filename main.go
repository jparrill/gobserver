package main

import (
	"context"
	"fmt"

	"github.com/jparrill/gobserver/internal/cmd"
)

var ctx context.Context
var cfg cmd.Config

func main() {

	// Initialize Context
	ctx = context.Background()

	// Recover the configuration
	cfg = cmd.RecoverConfig()

	// Initialize the logger
	logger := cmd.InitLogger(ctx)
	logger.Info("Loading Configuration")

	fmt.Printf("%T, %v", cfg, cfg.DB.DBName)

}
