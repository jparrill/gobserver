package database

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GOBSqlite struct {
	Kind string
	Dsn  string
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
		cmd.MainLogger.Panic("Engine MySQL not implemented")

	// PostgreSQL engine needs an IP:PORT, then perform the connection against the server
	case "postgres":
		cmd.MainLogger.Panic("Engine Postgres not implemented")

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

	for i, v := range fixtures {
		cmd.MainLogger.Sugar().Infof("Creating Asset %d: %v", i, v)
		db.Create(&entities.Organization{
			Name: v.OrgName,
		})

		db.Create(&entities.MLModel{
			Name:           v.MlmodelName,
			OrganizationID: uint(v.OrgId),
			Successes:      uint(v.MlmodelSuccesses),
			Fails:          uint(v.MlmodelFails),
		})
	}

}
