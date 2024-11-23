package cmd

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/helpers"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	token := helpers.ExtractBearerToken(authHeader)
	if token == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - invalid bearer token format")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrInvalidAuthorizationFormat))
		ctx.Abort()
		return
	}

	_, err := d.UserSessionRepository.FindUserSessionByToken(ctx, token)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrFindUserSessionByToken) {
			helpers.Logger.Error("middleware::MiddlewareValidateAuth - user session not found by token")
			ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrFindUserSessionByToken))
			ctx.Abort()
			return
		}

		helpers.Logger.Error("middleware::MiddlewareValidateAuth - failed to find user session by token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - failed to validate token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - token is already expired")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenExpired))
		ctx.Abort()
		return
	}

	ctx.Set(constants.TokenTypeAccess, claims)

	ctx.Next()
}

func (d *Dependency) MiddlewareRefreshToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	token := helpers.ExtractBearerToken(authHeader)
	if token == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - invalid bearer token format")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrInvalidAuthorizationFormat))
		ctx.Abort()
		return
	}

	_, err := d.UserSessionRepository.FindUserSessionByRefreshToken(ctx, token)
	if err != nil {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - failed to find user session by refresh token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - failed to validate refresh token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - refresh token is already expired")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenExpired))
		ctx.Abort()
		return
	}

	ctx.Set(constants.TokenTypeAccess, claims)

	ctx.Next()
}
