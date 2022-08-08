// Entities package contains all the models for the Application, which will be highly attached to Gorm.
// All entities which contains the gorm.Model, by default will include ID, CreatedAt, UpdatedAt and DeletedAt fields
package entities

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// Organization struct contains the Name and the inherited gorm.Model fields
type Organization struct {
	gorm.Model
	Name string `gorm:"Unique" json:"name" db:"name"`
}

// TableName function returns the TableName
func (org *Organization) TableName() string {
	return "organizations"
}

// ToJson returns string formatted organization struct
func (org Organization) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\n", org.ID, org.Name)
}

// ToJson returns JSON formatted organization struct
func (org Organization) ToJson() ([]byte, error) {
	return json.Marshal(&org)
}
