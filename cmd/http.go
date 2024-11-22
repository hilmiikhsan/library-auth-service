package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/api"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	services "github.com/hilmiikhsan/library-auth-service/internal/services/health_check"
)

func ServeHTTP() {
	dependency := dependencyInject()

	router := gin.Default()

	router.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
	}
}
