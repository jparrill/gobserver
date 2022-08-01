package entities

import (
	"fmt"

	"gorm.io/gorm"
)

// History struct contains OrganizationID (ForeignKey), MLModelID (ForeignKey), Success (Bool) and the inherited gorm.Model fields
type History struct {
	gorm.Model
	OrganizationID uint `db:"organizationid" json:"organizationid"`
	Organization   Organization
	MLModelID      uint `db:"mlmodelid" json:"mlmodelid"`
	MLModel        MLModel
	Success        bool `db:"success" json:"-"`
}

func (h *History) TableName() string {
	return "history"
}

func (h History) ToString() string {
	return fmt.Sprintf("id: %d\norg_id: %d\nmlmodel_id: %d\nsuccess: %v", h.ID, h.OrganizationID, h.MLModelID, h.Success)
}
