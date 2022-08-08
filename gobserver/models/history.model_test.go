package models_test

import (
	"testing"

	"github.com/jparrill/gobserver/models"
	"github.com/stretchr/testify/assert"
)

func TestHistoryModel(t *testing.T) {
	// Setup the env
	setUp()

	t.Run("TestHMFindAll", func(t *testing.T) {
		var hm models.HistoryModel

		hms, err := hm.FindAll()
		if err != nil {
			recover()
		}

		assert.True(t, len(hms) == 0)
	})
	// Tear env down
	tearDown()
}
