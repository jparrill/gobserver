// Entities package contains all the models for the Application, which will be highly attached to Gorm.
// All entities which contains the gorm.Model, by default will include ID, CreatedAt, UpdatedAt and DeletedAt fields
package entities

import (
	"fmt"
)

// Organization struct contains the Name and the inherited gorm.Model fields
type Organization struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"Unique"`
}

func (org *Organization) TableName() string {
	return "organizations"
}

func (org Organization) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\n", org.ID, org.Name)
}
