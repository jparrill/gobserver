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

	var orgModel models.OrganizationModel
	orgs := orgModel.FindAll()
	for _, org := range orgs {
		o, _ := org.ToJson()
		fmt.Println(string(o))
		fmt.Println("--------------")
	}

	//var mlModModel models.MLModModel
	//mlmodels := mlModModel.FindAll()
	//for _, ml := range mlmodels {
	//	fmt.Println(ml.ToString())
	//	fmt.Println("--------------")
	//}

	//orgModel.CreateOrg("client3")

}
