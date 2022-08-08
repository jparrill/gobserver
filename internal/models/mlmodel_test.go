package models_test

import (
	"log"
	"testing"

	"github.com/jparrill/gobserver/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMLModModel(t *testing.T) {
	// Setup the env
	setUp()

	t.Run("TestMLMFindAll", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindAll()
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].ID == 1)
		assert.True(t, mlmms[0].Name == "Decision tree")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 0)
	})

	t.Run("TestMLMFindAllWithFails", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindAllWithFails()
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].OrganizationID == 2)
		assert.True(t, mlmms[0].Name == "Convex Haul")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 9)
	})

	t.Run("TestMLMFindInOrgNoFails", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindInOrgNoFails(1)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].ID == 1)
		assert.True(t, mlmms[0].Name == "Decision tree")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 0)
	})

	t.Run("TestMLMFindAllNoFails", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindAllNoFails()
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].ID == 1)
		assert.True(t, mlmms[0].Name == "Decision tree")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 0)
	})

	t.Run("TestMLMFindInOrgWithFails", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindInOrgWithFails(2)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].OrganizationID == 2)
		assert.True(t, mlmms[0].Name == "Convex Haul")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 9)
	})

	t.Run("TestMLMFindAllInOrg", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindAllInOrg(1)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].OrganizationID == 1)
		assert.True(t, mlmms[0].Name == "Decision tree")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 0)
	})

	t.Run("TestMLMFindByName", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindByName("Convex Haul")
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].OrganizationID == 2)
		assert.True(t, mlmms[0].Name == "Convex Haul")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 9)
	})

	t.Run("TestMLMFindByNameInOrg", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindByNameInOrg("Convex Haul", 2)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms[0].OrganizationID == 2)
		assert.True(t, mlmms[0].Name == "Convex Haul")
		assert.True(t, mlmms[0].Successes == 0)
		assert.True(t, mlmms[0].Fails == 9)
	})

	t.Run("TestMLMFindById", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.FindById(1)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms.OrganizationID == 1)
		assert.True(t, mlmms.Name == "Decision tree")
		assert.True(t, mlmms.Successes == 0)
		assert.True(t, mlmms.Fails == 0)
	})

	t.Run("TestMLMCreateMLModel", func(t *testing.T) {
		var mlmm models.MLModModel

		mlmms, err := mlmm.CreateMLModel("Pepe Viyuela", 1)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, mlmms.OrganizationID == 1)
		assert.True(t, mlmms.Name == "Pepe Viyuela")
		assert.True(t, mlmms.Successes == 0)
		assert.True(t, mlmms.Fails == 0)
	})

	// Tear env down
	tearDown()
}
