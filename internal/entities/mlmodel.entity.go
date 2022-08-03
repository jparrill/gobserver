package entities

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// MLModel struct contains Name, OrganizationID (ForeignKey), Successes, Fails and the inherited gorm.Model fields
type MLModel struct {
	gorm.Model
	Name           string `db:"name" json:"name"`
	Organization   Organization
	OrganizationID uint
	Successes      uint `db:"successes" json:"successes"`
	Fails          uint `db:"fails" json:"fails"`
}

// TableName function returns the TableName
func (mlmodel *MLModel) TableName() string {
	return "mlmodels"
}

// ToJson returns string formatted mlmodel struct
func (mlmodel MLModel) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\norgId: %d\nsuccesses: %d\nfails: %d\n", mlmodel.ID, mlmodel.Name, mlmodel.OrganizationID, mlmodel.Successes, mlmodel.Fails)
}

// ToJson returns JSON formatted mlmodel struct
func (mlmodel MLModel) ToJson() ([]byte, error) {
	return json.Marshal(&mlmodel)
}
