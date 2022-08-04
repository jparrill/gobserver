package main

import (
	"context"
	"fmt"

	"github.com/jparrill/gobserver/internal/cmd"
	"github.com/jparrill/gobserver/internal/database"
	"github.com/jparrill/gobserver/internal/models"
)

var ctx context.Context

func main() {

	// Initialize Context
	ctx = context.Background()

	// Recover the configuration
	cmd.RecoverConfig()

	// Initialize the logger
	cmd.InitLogger()
	cmd.MainLogger.Info("Initializing DDBB")

	// Recovering DB handler from initialization
	db := database.Initialize(cmd.CFG.DB.DBType)

	database.Prepopulate(db)

	var mlmodsmod models.MLModModel
	mlmods, _ := mlmodsmod.FindAll()
	for _, mlmod := range mlmods {
		ml, _ := mlmod.ToJson()
		fmt.Println(string(ml))
		fmt.Println("--------------")
	}

	//var mlModModel models.MLModModel
	//mlmodels := mlModModel.FindAll()
	//for _, ml := range mlmodels {
	//	fmt.Println(ml.ToString())
	//	fmt.Println("--------------")
	//}

}
