package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/dto"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	"github.com/hilmiikhsan/library-auth-service/internal/validator"
)

type AuthHandler struct {
	AuthService interfaces.IAuthService
	Validator   *validator.Validator
}

func (api *AuthHandler) Register(ctx *gin.Context) {
	var (
		req = new(dto.RegisterRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::Register - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::Register - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	res, err := api.AuthService.Register(ctx.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUsernameAlreadyRegistered) {
			helpers.Logger.Error("handler::Register - Username already registered : ", err)
			ctx.JSON(http.StatusConflict, helpers.Error(constants.ErrUsernameAlreadyRegistered))
			return
		}

		helpers.Logger.Error("handler::Register - Failed to register user : ", err)
		code, errs := helpers.Errors[error](err)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}

func (api *AuthHandler) Login(ctx *gin.Context) {
	var (
		req = new(dto.LoginRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::Login - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::Login - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	res, err := api.AuthService.Login(ctx.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUsernameOrPasswordIsIncorrect) {
			helpers.Logger.Error("handler::Login - Username or password is incorrect : ", err)
			ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrUsernameOrPasswordIsIncorrect))
			return
		}

		helpers.Logger.Error("handler::Login - Failed to login user : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}

func (api *AuthHandler) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("handler::Logout - Authorization header is empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenIsEmpty))
		return
	}

	token := helpers.ExtractBearerToken(authHeader)

	err := api.AuthService.Logout(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("handler::Logout - Failed to logout user : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(nil, ""))
}
