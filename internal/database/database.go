package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/entities"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GOBSqlite struct {
	Kind string
	Dsn  string
}

type GOBMysql struct {
	Kind   string
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
	Dsn    string
}

type GOBPostgreSQL struct {
	Kind       string
	DBUser     string
	DBPass     string
	DBName     string
	DBHost     string
	DBPort     string
	DBSSL      string
	DBTimeZone string
	Dsn        string
}

type Object struct {
	OrgId            int
	OrgName          string
	MlmodelName      string
	MlmodelSuccesses int
	MlmodelFails     int
}

type Setup interface {
	Connect() *gorm.DB
	//Close()
}

func (sq GOBSqlite) Connect() *gorm.DB {
	cmd.MainLogger.Sugar().Debugf("---DEBUG---> DDBB Kind: %s DSN: %s", sq.Kind, sq.Dsn)
	db, err := gorm.Open(sqlite.Open(sq.Dsn), &gorm.Config{})
	if err != nil {
		cmd.MainLogger.Sugar().Panicf("Error recovering DB File in filepath: %s", sq.Dsn)
	}
	db.Exec("PRAGMA foreign_keys = ON;")
	return db
}

func (my GOBMysql) Connect() *gorm.DB {
	cmd.MainLogger.Sugar().Debugf("---DEBUG---> DDBB Kind: %s DSN: %s", my.Kind, my.Dsn)
	db, err := gorm.Open(mysql.Open(my.Dsn), &gorm.Config{})
	if err != nil {
		cmd.MainLogger.Sugar().Panicf("Error connecting MySQL DDBB: %s", my.Dsn)
	}
	return db
}

func (pg GOBPostgreSQL) Connect() *gorm.DB {
	cmd.MainLogger.Sugar().Debugf("---DEBUG---> DDBB Kind: %s DSN: %s", pg.Kind, pg.Dsn)
	db, err := gorm.Open(postgres.Open(pg.Dsn), &gorm.Config{})
	if err != nil {
		cmd.MainLogger.Sugar().Panicf("Error connecting PostgreSQL DDBB: %s", pg.Dsn)
	}
	return db
}

// Migrate function executes a DDBB AutoMigrate to precreate the tables, indexes, etc...
// For more info check "go doc gorm.DB.Automigrate"
func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&entities.Organization{},
		&entities.MLModel{},
		&entities.History{},
	); err != nil {
		cmd.MainLogger.Panic("Unable autoMigrateDB - " + err.Error())
	}
}

// Connector function receives the Setup interface, which matches the receiver with every
// kind of DDBB based on their type. Then the underhood function tries to connect it with gorm engine
// and it returns the handler.
func Connector(db Setup) *gorm.DB {
	return db.Connect()
}

// Initialize function receives a driver string based on the implemented DDBB types.
// It identifies the driver creating a type and trying to connect with the DDBB.
// Depending on the DDBB type, the connection method could change.
func Initialize(driver string) *gorm.DB {
	var db *gorm.DB

	db = GetDB(driver)
	Migrate(db)
	cmd.MainLogger.Sugar().Infof(`DDBB Initialized with "%s" driver`, driver)

	return db
}

// GetDB function returns a DDBB handler
func GetDB(driver string) *gorm.DB {
	var db *gorm.DB

	switch driver {
	// SQLite engine only needs to open an OS File
	case "sqlite":
		sqlite_db := GOBSqlite{
			Kind: driver,
			Dsn:  string(cmd.CFG.TMPFolder + cmd.CFG.DB.DBName),
		}
		db = Connector(sqlite_db)

	// MySQL engine needs an IP:PORT, then a ping and then perform the connection against the server
	case "mysql":
		mysql_db := GOBMysql{
			Kind:   driver,
			DBName: cmd.CFG.DB.DBName,
			DBUser: cmd.CFG.DB.DBUser,
			DBPass: cmd.CFG.DB.DBPass,
			DBHost: cmd.CFG.DB.DBHost,
			DBPort: cmd.CFG.DB.DBPort,
			Dsn: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cmd.CFG.DB.DBUser,
				cmd.CFG.DB.DBPass,
				cmd.CFG.DB.DBHost,
				cmd.CFG.DB.DBPort,
				cmd.CFG.DB.DBName,
			),
		}
		db = Connector(mysql_db)

	// PostgreSQL engine needs an IP:PORT, then perform the connection against the server
	case "postgres":
		pgsql_db := GOBPostgreSQL{
			Kind:       driver,
			DBName:     cmd.CFG.DB.DBName,
			DBUser:     cmd.CFG.DB.DBUser,
			DBPass:     cmd.CFG.DB.DBPass,
			DBHost:     cmd.CFG.DB.DBHost,
			DBPort:     cmd.CFG.DB.DBPort,
			DBSSL:      cmd.CFG.DB.DBSSL,
			DBTimeZone: cmd.CFG.DB.DBTimeZone,
			Dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
				cmd.CFG.DB.DBHost,
				cmd.CFG.DB.DBUser,
				cmd.CFG.DB.DBPass,
				cmd.CFG.DB.DBName,
				cmd.CFG.DB.DBPort,
				cmd.CFG.DB.DBSSL,
				cmd.CFG.DB.DBTimeZone,
			),
		}
		db = Connector(pgsql_db)

	}

	return db
}

func Prepopulate(db *gorm.DB) {

	var fixtures []Object

	fixturesFile, err := ioutil.ReadFile("fixtures/prepopulate_db.json")
	if err != nil {
		cmd.MainLogger.Sugar().Errorf("Error converting JSON file to []byte: %s", err)
	}

	err = json.Unmarshal([]byte(fixturesFile), &fixtures)
	if err != nil {
		cmd.MainLogger.Sugar().Errorf("Error unmarshalling fixtures file: %s", err)
	}

	var org entities.Organization
	var mlmodel entities.MLModel

	for i, v := range fixtures {
		var q entities.Organization
		// Insert Organization
		cmd.MainLogger.Sugar().Infof("Creating Asset %d: %v", i, v)

		// Check the OrgID from a query
		db.Table("organizations").Where("Name = ?", v.OrgName).Find(&q)

		// If the org not exists, create it
		if q.Name == "" {
			org = entities.Organization{
				Name: v.OrgName,
			}
			db.Create(&org)
		}

		// If the org did not exists, the id is empty, so we recover again the resource
		if q.ID == 0 {
			db.Table("organizations").Where("Name = ?", v.OrgName).Find(&q)
		}
		// Insert ML Model into Into Organization
		mlmodel = entities.MLModel{
			Name:           v.MlmodelName,
			OrganizationID: uint(q.ID),
			Successes:      uint(v.MlmodelSuccesses),
			Fails:          uint(v.MlmodelFails),
		}
		db.Create(&mlmodel)
	}

}
