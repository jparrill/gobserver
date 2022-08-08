package config_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jparrill/gobserver/config"
	"github.com/stretchr/testify/assert"
)

func TestLogFileAfterInit(t *testing.T) {
	configFile := filepath.Join("testdata", "config.sqlite.json")
	config.RecoverConfig(configFile)
	config.InitLogger()

	_, err := os.Stat(config.CFG.TMPFolder)
	if err != nil {
		log.Panicf("Error openning the %s logfile to store logs\n", config.CFG.Log.LogPath)
	}
}

func TestLogMessage(t *testing.T) {
	configFile := filepath.Join("testdata", "config.sqlite.json")
	config.RecoverConfig(configFile)
	config.InitLogger()

	config.MainLogger.Info("TEST Info Message")

	b, err := ioutil.ReadFile(config.CFG.Log.LogPath)
	if err != nil {
		log.Panicf("Logfile does not exists in path %s: %v", config.CFG.Log.LogPath, err)
	}
	assert.True(t, strings.Contains(string(b), "TEST Info Message"))
}
