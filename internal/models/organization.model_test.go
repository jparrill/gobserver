package models_test

import (
	"log"
	"testing"

	"github.com/jparrill/gobserver/internal/models"
	"github.com/stretchr/testify/assert"
)

//OM == Organization Model

func TestOrganizationModel(t *testing.T) {
	// Setup the env
	setUp()

	t.Run("TestOMFindAll", func(t *testing.T) {
		var om models.OrganizationModel

		orgs_t, err := om.FindAll()
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, orgs_t[0].ID == 1)
		assert.True(t, orgs_t[0].Name == "client1")
	})

	t.Run("TestOMOrgExists", func(t *testing.T) {
		var om models.OrganizationModel

		org_exists := om.OrgExists("client1")
		org_not_exists := om.OrgExists("client8")
		assert.True(t, org_exists)
		assert.False(t, org_not_exists)
	})

	t.Run("TestOMFindByName", func(t *testing.T) {
		var om models.OrganizationModel

		org, err := om.FindByName("client1")
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, org.Name == "client1")
	})

	t.Run("TestOMFindById", func(t *testing.T) {
		var om models.OrganizationModel

		org, err := om.FindById(1)
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, org.Name == "client1")
	})

	t.Run("TestOMCreateOrg", func(t *testing.T) {
		var om models.OrganizationModel

		org, err := om.CreateOrg("testOrg")
		if err != nil {
			log.Panicf("Error recovering data from DDBB: %v", err)
		}
		assert.True(t, org.Name == "testOrg")

	})

	// Tear env down
	tearDown()
}
