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

func (api *AuthHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
		req = new(dto.RegisterRequest)
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("handler::Register - Failed to bind request : ", err)
		c.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		log.Error("handler::Register - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		c.JSON(code, helpers.Error(errs))
		return
	}

	res, err := api.AuthService.Register(c, req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUsernameAlreadyRegistered) {
			log.Error("handler::Register - Username already registered : ", err)
			c.JSON(http.StatusConflict, helpers.Error(constants.ErrUsernameAlreadyRegistered))
			return
		}

		log.Error("handler::Register - Failed to register user : ", err)
		code, errs := helpers.Errors[error](err)
		c.JSON(code, helpers.Error(errs))
		return
	}

	c.JSON(http.StatusOK, helpers.Success(res, ""))
}
