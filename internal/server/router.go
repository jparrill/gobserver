package server

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jparrill/gobserver/internal/config"
	"github.com/jparrill/gobserver/internal/controllers"
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
			orgGroup.GET("/:id", org.Retrieve)
		}
	}
	return router

}
