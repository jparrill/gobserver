// Model package contains all the models for the Application, which will be highly attached to Gorm.
// All models which contains the gorm.Model, by default will include ID, CreatedAt, UpdatedAt and DeletedAt fields

package model

import (
	"gorm.io/gorm"
)

// Organization struct contains the Name and the inherited gorm.Model fields
type Organization struct {
	gorm.Model
	Name string `db:"name" json:"name"`
}

// MLModel struct contains Name, OrganizationID (ForeignKey), Successes, Fails and the inherited gorm.Model fields
type MLModel struct {
	gorm.Model
	Name           string `db:"name" json:"name"`
	OrganizationID uint   `db:"organizationid" json:"organizationid"`
	Successes      uint   `db:"successes" json:"successes"`
	Fails          uint   `db:"fails" json:"fails"`
}

// History struct contains OrganizationID (ForeignKey), MLModelID (ForeignKey), Success (Bool) and the inherited gorm.Model fields
type History struct {
	gorm.Model
	OrganizationID uint `db:"organizationid" json:"organizationid"`
	MLModelID      uint `db:"mlmodelid" json:"mlmodelid"`
	Success        bool `db:"sucess" json:"-"`
}
