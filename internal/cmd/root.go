package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

const (
	basePath = "/tmp/gobserver"
)

//InitLogger function initializes the Logger engine using Zap as a base
func InitLogger() *zap.Logger {
	// TODO: This Function receives the Context, we need to store the logger inside
	// in order to use it among the whole app
	_, err := os.Stat(basePath)
	if err != nil {
		fmt.Println("BasePath for Log entries does not exists, creating...")
		os.Mkdir(basePath, 0754)
	}

	rawJSON := []byte(`{
	  "level": "` + CFG.Log.LogLevel + `",
	  "encoding": "json",
	  "outputPaths": ["stdout", "` + basePath + `/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "uppercase"
	  }
	}`)

	var loggerConfig zap.Config

	if err := json.Unmarshal(rawJSON, &loggerConfig); err != nil {
		panic(err)
	}
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("Logger construction succeeded")

	return logger
}
