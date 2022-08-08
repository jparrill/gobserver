package config_test

import (
	"path/filepath"
	"testing"

	"github.com/jparrill/gobserver/internal/config"
	"github.com/stretchr/testify/assert"
)

// JSON Tests
func TestRecoverJSONConfigMySQL(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.mysql.json")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "mysql")
}

func TestRecoverJSONConfigSqlite(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.sqlite.json")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "sqlite")
}

func TestRecoverJSONConfigPgSQL(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.pgsql.json")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "postgres")
}

// YAML Tests
func TestRecoverYamlConfigMySQL(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.mysql.yaml")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "mysql")
}

func TestRecoverYamlConfigSqlite(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.sqlite.yaml")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "sqlite")
}

func TestRecoverYamlConfigPgSQL(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.pgsql.yaml")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "postgres")
}

// TOML Tests
func TestRecoverTOMLConfigMySQL(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.mysql.toml")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "mysql")
}

func TestRecoverTOMLConfigSqlite(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.sqlite.toml")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "sqlite")
}

func TestRecoverTOMLConfigPgSQL(t *testing.T) {
	// Locate the current directory
	configFile := filepath.Join("testdata", "config.pgsql.toml")
	config.RecoverConfig(configFile)
	assert.Equal(t, config.CFG.DB.DBType, "postgres")
}
