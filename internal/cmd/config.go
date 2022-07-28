package cmd

import (
	"context"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct holds the fields for global config
type Config struct {
	// Path to
	TmpFolder string `yaml:"tmpfolder", json:"tmpfolder"`
	// DB realted config
	DB struct {
		// DBName is a Filename in sqlite or DDBB Name in case of Mysql or Postgres
		DBName string `yaml:"dbname", json:"dbtype"`
		// DBType: "sqlite"|"mysql"|"postgres"
		DBType string `yaml:"dbtype", json:"dbtype"`
		// DBUser
		DBUser string `yaml:"dbuser", json:"dbuser"`
		// DBPass
		DBPass string `yaml:"dbpass", json:"dbpass"`
	} `yaml:"db", json:"db"`
	// Log related config
	Log struct {
		LogPath  string `yaml:"logpath", json:"logpath"`
		LogLevel string `yaml:"loglevel", json:"loglevel"`
	} `yaml:"log", json:"log"`
}

type Manage interface {
	Recover(Config) Config
}

func (cfg *Config) Recover() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("ERROR: Early error recovering the config file: %s", "config.yaml")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func RecoverConfig(ctx context.Context) Config {
	var cfg Config
	err := cfg.Recover()
	if err != nil {
		panic("Error decoding Yaml")
	}
	return cfg
}
