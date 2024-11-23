package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/helpers"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	_, err := d.UserSessionRepository.FindUserSessionByToken(ctx, auth)
	if err != nil {
		log.Println("failed to find user session by token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}
}
