package database

import (
	"fmt"

	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var org model.Organization

type DDBB struct{}
type gobSqlite struct {
	kind string
	dsn  string
}

type Setup interface {
	//InitDB()
	Connect() *gorm.DB
	Close()
}

/*func (sq sqlite) InitDB() {
	_, err := os.Stat(CFG.)
	if err != nil {
		fmt.Println("BasePath for Log entries does not exists, creating...")
		os.Mkdir(CFG.TMPFolder, 0754)
	}
}
*/

func (sq gobSqlite) Connect() *gorm.DB {
	cmd.MainLogger.Debug(sq.dsn)
	db, err := gorm.Open(sqlite.Open(sq.dsn), &gorm.Config{})
	if err != nil {
		cmd.MainLogger.Error("Error recovering DB Filename")
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Organization{}, &model.MLModel{}, &model.History{})
}

func Identify(driver string) {
	var query string

	switch driver {
	case "sqlite":
		sqlite_db := gobSqlite{
			kind: driver,
			dsn:  string(cmd.CFG.TMPFolder + cmd.CFG.DB.DBName),
		}
		db := sqlite_db.Connect()
		db.AutoMigrate(&model.Organization{}, &model.MLModel{}, &model.History{})
		db.Create(&model.Organization{Name: "org1"})
		result := map[string]interface{}{}
		db.First(&org)
		db.Table("organizations").Take(&result)
		for k, v := range result {
			query += fmt.Sprintf("%v: %v\n", k, v)
			//fmt.Printf("ID: %d,Name: %s\n", result["id"], result["name"])
		}
		cmd.MainLogger.Info(query)

	}

}
