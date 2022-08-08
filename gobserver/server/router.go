package server

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jparrill/gobserver/config"
	"github.com/jparrill/gobserver/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	a := ginzap.Ginzap(config.MainLogger, time.RFC3339, true)
	router.Use(a)
	router.Use(ginzap.RecoveryWithZap(config.MainLogger, true))

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		orgGroup := v1.Group("org")
		{
			org := new(controllers.OrganizationController)
			orgGroup.GET("/name/:name", org.Retrieve)
			orgGroup.GET("/id/:id", org.Retrieve)
			orgGroup.POST("/add", org.Create)
		}
		mlmodGroup := v1.Group("ml")
		{
			mlmod := new(controllers.MLModelController)
			mlmodGroup.GET("/id/:orgId/:id", mlmod.RetrieveId)
			mlmodGroup.GET("/nofails/:orgId", mlmod.RetrieveModInOrgNoFails)
			mlmodGroup.GET("/nofails/all", mlmod.RetrieveModNoFails)
			mlmodGroup.POST("/add", mlmod.Create)
			mlmodGroup.POST("/addBulk", mlmod.CreateBulk)
		}
	}
	return router

}
