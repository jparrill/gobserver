package database

import (
	"fmt"

	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GOBSqlite struct {
	Kind string
	Dsn  string
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
	return db
}

// Migrate function executes a DDBB AutoMigrate to precreate the tables, indexes, etc...
// For more info check "go doc gorm.DB.Automigrate"
func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.Organization{},
		&model.MLModel{},
		&model.History{},
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
		cmd.MainLogger.Panic("Engine MySQL not implemented")

	// PostgreSQL engine needs an IP:PORT, then perform the connection against the server
	case "postgres":
		cmd.MainLogger.Panic("Engine Postgres not implemented")

	}
	Migrate(db)
	cmd.MainLogger.Sugar().Infof(`DDBB Initialized with "%s" driver`, driver)

	return db
}

func Prepopulate(db *gorm.DB) {

	db.Create(&model.Organization{Name: "org1"})

}

func Query(db *gorm.DB) {

	var query string
	var org model.Organization

	result := map[string]interface{}{}
	db.First(&org)
	db.Table("organizations").Take(&result)
	for k, v := range result {
		query += fmt.Sprintf("%v: %v\n", k, v)
		//fmt.Printf("ID: %d,Name: %s\n", result["id"], result["name"])
	}
	cmd.MainLogger.Info(query)
}
