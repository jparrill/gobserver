package cmd

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var cfg Config

type JSONCfgFile Config
type YAMLCfgFile Config
type TOMLCfgFile Config

// Config struct holds the fields for global config
type Config struct {
	// DB realted config
	DB struct {
		// DBName is a Filename in sqlite or DDBB Name in case of Mysql or Postgres
		DBName string `yaml:"dbname" json:"dbname"`
		// DBType: "sqlite"|"mysql"|"postgres"
		DBType string `yaml:"dbtype" json:"dbtype"`
		// DBUser
		DBUser string `yaml:"dbuser" json:"dbuser"`
		// DBPass
		DBPass string `yaml:"dbpass" json:"dbpass"`
	} `yaml:"db" json:"db"`
	// Log related config
	Log struct {
		LogPath  string `yaml:"logpath" json:"logpath"`
		LogLevel string `yaml:"loglevel" json:"loglevel"`
	} `yaml:"log" json:"log"`
}

// Manage interface gives the methods to cover different source files
// like YAML, TOML, JSON. First we create a new type from config to associate it
// to a method of Recover function, then develop the driver and put the
// logic on RecoverConfig function to select the correct type and driver.
type Manage interface {
	Recover() error
}

func (cfg *YAMLCfgFile) Recover() error {
	cf := "config.yaml"
	f, err := os.Open(cf)
	if err != nil {
		log.Fatalf("ERROR: Early error recovering the config file: %s", cf)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Panic(err)
	}

	return nil
}
func (cfg *JSONCfgFile) Recover() error {
	// To be implemented
	cf := "config.json"
	fmt.Println(cf)
	return nil
}

func (cfg *TOMLCfgFile) Recover() error {
	// To be implemented
	cf := "config.toml"
	fmt.Println(cf)
	return nil
}

// RecoverConfig function will recover Config File from the repo's root folder
// it could be JSON, YAML or TOML.
func RecoverConfig() Config {

	var configFile Config

	if _, err := os.Stat("config.json"); err == nil {
		cfg := JSONCfgFile(configFile)
		err := cfg.Recover()
		if err != nil {
			panic("Error decoding JSON")
		}
		return Config(cfg)
	} else if _, err := os.Stat("config.yaml"); err == nil {
		cfg := YAMLCfgFile(configFile)
		err := cfg.Recover()
		if err != nil {
			panic("Error decoding YAML")
		}
		return Config(cfg)
	} else if _, err := os.Stat("config.toml"); err == nil {
		cfg := TOMLCfgFile(configFile)
		err := cfg.Recover()
		if err != nil {
			panic("Error decoding TOML")
		}
		return Config(cfg)
	} else {
		log.Panic(`Config File does not exists, it should format: "yaml", "json" or "toml" with name "config"`)
		return cfg
	}
}
