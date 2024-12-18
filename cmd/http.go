package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	authAPI "github.com/hilmiikhsan/library-auth-service/internal/api/auth"
	healthCheckAPI "github.com/hilmiikhsan/library-auth-service/internal/api/health_check"
	api "github.com/hilmiikhsan/library-auth-service/internal/api/token_validation"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	userRepository "github.com/hilmiikhsan/library-auth-service/internal/repository/user"
	userSessionRepository "github.com/hilmiikhsan/library-auth-service/internal/repository/user_session"
	authServices "github.com/hilmiikhsan/library-auth-service/internal/services/auth"
	healthCheckServices "github.com/hilmiikhsan/library-auth-service/internal/services/health_check"
	tokenValidationServices "github.com/hilmiikhsan/library-auth-service/internal/services/token_validation"
	"github.com/hilmiikhsan/library-auth-service/internal/validator"
	"github.com/sirupsen/logrus"
)

func ServeHTTP() {
	dependency := dependencyInject()

	router := gin.Default()

	router.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := router.Group("/user/v1")
	userV1.POST("/register", dependency.AuthAPI.Register)
	userV1.POST("/login", dependency.AuthAPI.Login)

	userV1WithAuth := userV1.Use()
	userV1WithAuth.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.AuthAPI.Logout)
	userV1WithAuth.PUT("/refresh-token", dependency.MiddlewareRefreshToken, dependency.AuthAPI.RefreshToken)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		helpers.Logger.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	Logger                *logrus.Logger
	UserRepository        interfaces.IUserRepository
	UserSessionRepository interfaces.IUserSessionRepository

	HealthcheckAPI     interfaces.IHealthcheckHandler
	AuthAPI            interfaces.IAuthHandler
	TokenValidationAPI *api.TokenValidationHandler
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

	userSessionRepo := &userSessionRepository.UserSessionRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	authSvc := &authServices.AuthService{
		UserRepo:        userRepo,
		UserSessionRepo: userSessionRepo,
		Logger:          helpers.Logger,
	}
	authAPI := &authAPI.AuthHandler{
		AuthService: authSvc,
		Validator:   validator,
	}

	tokenValidationSvc := &tokenValidationServices.TokenValidationService{
		UserSessionRepo: userSessionRepo,
		Logger:          helpers.Logger,
	}
	tokenValidationAPI := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationSvc,
	}

	return Dependency{
		Logger:                helpers.Logger,
		UserRepository:        userRepo,
		UserSessionRepository: userSessionRepo,
		HealthcheckAPI:        healthcheckAPI,
		AuthAPI:               authAPI,
		TokenValidationAPI:    tokenValidationAPI,
	}
}
