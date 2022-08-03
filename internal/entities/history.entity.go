package entities

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// History struct contains OrganizationID (ForeignKey), MLModelID (ForeignKey), Success (Bool) and the inherited gorm.Model fields
type History struct {
	gorm.Model
	Organization   Organization
	OrganizationID uint
	MLModel        MLModel
	MLModelID      uint
	Success        bool `db:"success" json:"-"`
}

// TableName function returns the TableName
func (h *History) TableName() string {
	return "history"
}

// ToJson returns string formatted history struct
func (h History) ToString() string {
	return fmt.Sprintf("id: %d\norg_id: %d\nmlmodel_id: %d\nsuccess: %v", h.ID, h.OrganizationID, h.MLModelID, h.Success)
}

// ToJson returns JSON formatted history struct
func (h History) ToJson() ([]byte, error) {
	return json.Marshal(&h)
}
