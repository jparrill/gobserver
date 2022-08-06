package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

var CFG Config

type JSONCfgFile Config
type YAMLCfgFile Config
type TOMLCfgFile Config

// Config struct holds the fields for global config
type Config struct {
	// TMPFolder is a temporary folder to hold all app assets
	TMPFolder string `yaml:"tmpfolder" json:"tmpfolder"`
	// AppPort where the application will bind in
	AppPort string `yaml:"appport" json:"appport"`
	// DB realted config
	DB struct {
		// DBName is a Filename in sqlite or DDBB Name in case of Mysql or Postgres
		// Relevant in all cases
		DBName string `yaml:"dbname" json:"dbname"`
		// DBType: "sqlite"|"mysql"|"postgres"
		// Relevant in all cases
		DBType string `yaml:"dbtype" json:"dbtype"`
		// DBUser "gobuser"
		// Only relevant in MySQL and PostgreSQL
		DBUser string `yaml:"dbuser" json:"dbuser"`
		// DBPass "P4$$w0rD"
		// Only relevant in MySQL and PostgreSQL
		DBPass string `yaml:"dbpass" json:"dbpass"`
		// DBHost "127.0.0.1"
		// Only relevant in MySQL and PostgreSQL
		DBHost string `yaml:"dbhost" json:"dbhost"`
		// DBPort "5432"
		// Only relevant in MySQL and PostgreSQL
		DBPort string `yaml:"dbport" json:"dbport"`
		// DBSSL "enable|disable"
		// Only relevant in MySQL and PostgreSQL
		DBSSL string `yaml:"dbssl" json:"dbssl"`
		// DBTimeZone "UTC" "Asia/Shanghai"
		// Only relevant in PostgreSQL
		DBTimeZone string `yaml:"dbtimezone" json:"dbtimezone"`
	} `yaml:"db" json:"db"`
	// Log related config
	Log struct {
		// Concanetated with TMPFolder
		LogPath string `yaml:"logpath" json:"logpath"`
		// Loglevel option can be these ones: debug|info|warn|error|panic|fatal.
		// For more info check gobserver/internal/cmd/root.go on the switch statement.
		// The log level are equivalent to zapcore.LevelEnabler type.
		// For more into check "go doc zapcore.LevelEnabler" or "go doc zapcore.DebugLevel"
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

// Recover function using the YAMLCfgFile as a driver, parses the YAML file loaded and
// injects the content into a Config struct and exposes it in CFG global var
func (cfg *YAMLCfgFile) Recover(configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("ERROR: Early error recovering the config file: %s", configPath)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("Error Decoding YAML Config file: %s", configPath)
	}

	return nil
}

// Recover function using the JSONCfgFile as a driver, parses the JSON file loaded and
// injects the content into a Config struct and exposes it in CFG global var
func (cfg *JSONCfgFile) Recover(configPath string) error {
	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("ERROR: Early error recovering the config file: %s", configPath)
	}

	err = json.Unmarshal([]byte(f), cfg)
	if err != nil {
		log.Fatalf("Error Unmarshalling JSON Config file: %s", configPath)
	}

	return nil
}

// Recover function using the TOMLCfgFile as a driver, parses the TOML file loaded and
// injects the content into a Config struct and exposes it in CFG global var
func (cfg *TOMLCfgFile) Recover(configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("ERROR: Early error recovering the config file: %s", configPath)
	}
	defer f.Close()

	decoder := toml.NewDecoder(f)
	_, err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("Error Decoding TOML Config file: %s", configPath)
	}

	return nil
}

// RecoverConfig function will recover Config File from the repo's root folder
// it could be JSON, YAML or TOML.
func RecoverConfig(configPath string) {

	if _, err := os.Stat(configPath); err != nil {
		log.Panicf(`Cannot open configFile in path: %s`, configPath)
	}

	var configFile Config
	// Get Basepath
	fileName := filepath.Base(configPath)

	// Get Extension
	ext := filepath.Ext(fileName)

	switch ext {
	case ".json":
		cfg := JSONCfgFile(configFile)
		err := cfg.Recover(configPath)
		if err != nil {
			log.Panicln("Error decoding JSON")
		}
		configFile = Config(cfg)

	case ".yaml", ".yml":
		cfg := YAMLCfgFile(configFile)
		err := cfg.Recover(configPath)
		if err != nil {
			log.Panicln("Error decoding YAML")
		}
		configFile = Config(cfg)

	case ".toml":
		cfg := TOMLCfgFile(configFile)
		err := cfg.Recover(configPath)
		if err != nil {
			log.Panicln("Error decoding TOML")
		}
		configFile = Config(cfg)
	default:
		log.Fatalf(`Error recovering Config file. The file %s does not contain any of the supported extension: "json|toml|yaml|yml"`, configPath)
	}

	CFG = configFile
}
