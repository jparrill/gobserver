package entities_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/jparrill/gobserver/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestHistEnt(t *testing.T) {

	he := entities.History{
		OrganizationID: 1,
		MLModelID:      1,
		Success:        true,
	}
	assert.True(t, he.ID == 0)
	assert.True(t, he.OrganizationID == 1)
	assert.True(t, he.MLModelID == 1)
	assert.True(t, he.Success == true)

}

func TestHETabName(t *testing.T) {

	he := entities.History{
		OrganizationID: 1,
		MLModelID:      1,
		Success:        true,
	}
	assert.True(t, he.TableName() == "history")
}

func TestHEToString(t *testing.T) {

	he := entities.History{
		OrganizationID: 1,
		MLModelID:      1,
		Success:        true,
	}
	assert.True(t, strings.Contains(he.ToString(), "1"))
	assert.True(t, strings.Contains(he.ToString(), "true"))
}

func TestHEToJSON(t *testing.T) {
	var JSONRes entities.History

	he := entities.History{
		OrganizationID: 1,
		MLModelID:      1,
		Success:        true,
	}
	bjson, err := he.ToJson()
	if err != nil {
		log.Panic("Error Converting History entity to JSON: ", err)
	}

	err = json.Unmarshal(bjson, &JSONRes)
	if err != nil {
		log.Panic("Error Unmarshalling: ", err)
	}
	assert.True(t, JSONRes.ID == 0)
	assert.True(t, JSONRes.OrganizationID == 1)
	assert.True(t, JSONRes.MLModelID == 1)
	assert.True(t, JSONRes.Success == true)
}
