package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	authAPI "github.com/hilmiikhsan/library-auth-service/internal/api/auth"
	healthCheckAPI "github.com/hilmiikhsan/library-auth-service/internal/api/health_check"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	userRepository "github.com/hilmiikhsan/library-auth-service/internal/repository/user"
	authServices "github.com/hilmiikhsan/library-auth-service/internal/services/auth"
	healthCheckServices "github.com/hilmiikhsan/library-auth-service/internal/services/health_check"
	"github.com/hilmiikhsan/library-auth-service/internal/validator"
	"github.com/sirupsen/logrus"
)

func ServeHTTP() {
	dependency := dependencyInject()

	router := gin.Default()

	router.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := router.Group("/user/v1")
	userV1.POST("/register", dependency.AuthAPI.Register)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	Logger         *logrus.Logger
	UserRepository interfaces.IUserRepository

	HealthcheckAPI interfaces.IHealthcheckHandler
	AuthAPI        interfaces.IAuthHandler
}

func dependencyInject() Dependency {
	helpers.SetupLogger()

	healthcheckSvc := &healthCheckServices.Healthcheck{}
	healthcheckAPI := &healthCheckAPI.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	userRepo := &userRepository.UserRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	authSvc := &authServices.AuthService{
		UserRepo: userRepo,
		Logger:   helpers.Logger,
	}
	authAPI := &authAPI.AuthHandler{
		AuthService: authSvc,
		Validator:   validator,
	}

	return Dependency{
		Logger:         helpers.Logger,
		UserRepository: userRepo,
		HealthcheckAPI: healthcheckAPI,
		AuthAPI:        authAPI,
	}
}
