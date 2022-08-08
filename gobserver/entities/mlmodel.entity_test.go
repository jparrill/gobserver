package entities_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/jparrill/gobserver/entities"
	"github.com/stretchr/testify/assert"
)

func TestMLEnt(t *testing.T) {

	mlme := entities.MLModel{
		Name:           "Convex Haul",
		OrganizationID: 1,
		Successes:      10,
		Fails:          5,
	}
	assert.True(t, mlme.ID == 0)
	assert.True(t, mlme.Name == "Convex Haul")
	assert.True(t, mlme.OrganizationID == 1)
	assert.True(t, mlme.Successes == 10)
	assert.True(t, mlme.Fails == 5)

}

func TestMLETabName(t *testing.T) {

	mlme := entities.MLModel{
		Name:           "Convex Haul",
		OrganizationID: 1,
		Successes:      10,
		Fails:          5,
	}
	assert.True(t, mlme.TableName() == "mlmodels")
}

func TestMLEToString(t *testing.T) {

	mlme := entities.MLModel{
		Name:           "Convex Haul",
		OrganizationID: 1,
		Successes:      10,
		Fails:          5,
	}
	assert.True(t, strings.Contains(mlme.ToString(), "Convex Haul"))
	assert.True(t, strings.Contains(mlme.ToString(), "0"))
	assert.True(t, strings.Contains(mlme.ToString(), "5"))
	assert.True(t, strings.Contains(mlme.ToString(), "1"))
	assert.True(t, strings.Contains(mlme.ToString(), "10"))
}

func TestMLEToJSON(t *testing.T) {
	var JSONRes entities.Organization

	mlme := entities.MLModel{
		Name:           "Convex Haul",
		OrganizationID: 1,
		Successes:      10,
		Fails:          5,
	}
	bjson, err := mlme.ToJson()
	if err != nil {
		log.Panic("Error Converting MLModel entity to JSON: ", err)
	}

	err = json.Unmarshal(bjson, &JSONRes)
	if err != nil {
		log.Panic("Error Unmarshalling: ", err)
	}
	assert.True(t, mlme.ID == 0)
	assert.True(t, mlme.Name == "Convex Haul")
	assert.True(t, mlme.OrganizationID == 1)
	assert.True(t, mlme.Successes == 10)
	assert.True(t, mlme.Fails == 5)
}
