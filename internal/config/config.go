package config

import (
	"fmt"
	"log"
	"os"

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

// Recover function using the JSONCfgFile as a driver, parses the JSON file loaded and
// injects the content into a Config struct and exposes it in CFG global var
func (cfg *JSONCfgFile) Recover() error {
	// To be implemented
	cf := "config.json"
	fmt.Println(cf)
	return nil
}

// Recover function using the TOMLCfgFile as a driver, parses the TOML file loaded and
// injects the content into a Config struct and exposes it in CFG global var
func (cfg *TOMLCfgFile) Recover() error {
	// To be implemented
	cf := "config.toml"
	fmt.Println(cf)
	return nil
}

// RecoverConfig function will recover Config File from the repo's root folder
// it could be JSON, YAML or TOML.
func RecoverConfig() {

	var configFile Config

	if _, err := os.Stat("config.json"); err == nil {
		cfg := JSONCfgFile(configFile)
		err := cfg.Recover()
		if err != nil {
			log.Panicln("Error decoding JSON")
		}
		configFile = Config(cfg)

	} else if _, err := os.Stat("config.yaml"); err == nil {
		cfg := YAMLCfgFile(configFile)
		err := cfg.Recover()
		if err != nil {
			log.Panicln("Error decoding YAML")
		}
		configFile = Config(cfg)

	} else if _, err := os.Stat("config.toml"); err == nil {
		cfg := TOMLCfgFile(configFile)
		err := cfg.Recover()
		if err != nil {
			log.Panicln("Error decoding TOML")
		}
		configFile = Config(cfg)

	} else {
		log.Panicln(`Config File does not exists, it should format: "yaml", "json" or "toml" with name "config"`)

	}

	CFG = configFile
}
