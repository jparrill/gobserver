package entities

import (
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

func (mlmodel *MLModel) TableName() string {
	return "mlmodels"
}

func (mlmodel MLModel) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\norgId: %d\nsuccesses: %d\nfails: %d\n", mlmodel.ID, mlmodel.Name, mlmodel.OrganizationID, mlmodel.Successes, mlmodel.Fails)
}
