package database

import (
	"log"
	"os"

	"github.com/jparrill/gormsample/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Setup interface {
	Connect() *gorm.DB
	Close()
}

func (sq sqlite) Connect() *gorm.DB {
	db_file, err := os.Environ("SQLITE_DB_NAME")
	if err != nil {
		log.Println("Error recovering DB Filename")
	}
	db_path := "/tmp/gormsample/" + db_file
	return gorm.Open(sqlite.Open(db_file), &gorm.Config{})
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Organization{}, &model.MLModel{}, &model.History{})
}
