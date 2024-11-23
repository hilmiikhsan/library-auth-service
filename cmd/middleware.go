package cmd

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/helpers"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Error("authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	_, err := d.UserSessionRepository.FindUserSessionByToken(ctx, auth)
	if err != nil {
		log.Error("failed to find user session by token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Error("failed to validate token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		log.Error("token is already expired")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenExpired))
		ctx.Abort()
		return
	}

	ctx.Set(constants.TokenTypeAccess, claims)

	ctx.Next()
}
