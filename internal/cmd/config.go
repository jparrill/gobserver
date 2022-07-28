package cmd

// Config struct holds the fields for global config
type Config struct {
	TmpFolder string `yaml:"tmpfolder"`
	// DB realted config
	DB struct {
		// DBName is a Filename in sqlite or DDBB Name in case of Mysql or Postgres
		DBName string `yaml:"dbname"`
		// DBPath
		DBPath string `yaml:"dbpath"`
		// DBType: "sqlite"|"mysql"|"postgres"
		DBType string `yaml:"dbtype"`
		// DBUser
		DBUser string `yaml:"dbuser"`
		// DBPass
		DBPass string `yaml:"pass"`
	} `yaml:"db"`
	// Log related config
	Log struct {
		LogPath  string `yaml:"logpath"`
		LogLevel string `yaml:"loglevel"`
	} `yaml:"log"`
}

var cfg Config

func RecoverConfig() {

}
