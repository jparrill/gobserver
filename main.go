package main

import (
	"github.com/jparrill/gormsample/internal/cmd"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {

	// Recover the configuration
	config = cmd.RecoverConfig()
	// Initialize the logger
	logger = cmd.InitLogger()
}
