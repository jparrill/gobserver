package entities_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/jparrill/gobserver/entities"
	"github.com/stretchr/testify/assert"
)

func TestOrgEnt(t *testing.T) {

	oe := entities.Organization{Name: "Paul"}
	assert.True(t, oe.Name == "Paul")
	assert.True(t, oe.ID == 0)

}

func TestOETabName(t *testing.T) {

	oe := entities.Organization{
		Name: "Paul",
	}
	assert.True(t, oe.TableName() == "organizations")
}

func TestOEToString(t *testing.T) {

	oe := entities.Organization{Name: "Paul"}
	assert.True(t, strings.Contains(oe.ToString(), "Paul"))
	assert.True(t, strings.Contains(oe.ToString(), "0"))
}

func TestOEToJSON(t *testing.T) {
	var JSONRes entities.Organization

	oe := entities.Organization{Name: "Paul"}
	bjson, err := oe.ToJson()
	if err != nil {
		log.Panic("Error Converting Org entity to JSON: ", err)
	}

	err = json.Unmarshal(bjson, &JSONRes)
	if err != nil {
		log.Panic("Error Unmarshalling: ", err)
	}
	assert.True(t, JSONRes.Name == "Paul")
	assert.True(t, JSONRes.ID == 0)
}
