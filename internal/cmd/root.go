package cmd

import (
	"encoding/json"

	"go.uber.org/zap"
)

//initLogger function initializes the Logger engine using Zap as a base
func InitLogger() *zap.Logger {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "uppercase"
	  }
	}`)

	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("Logger construction succeeded")

	return logger
}
