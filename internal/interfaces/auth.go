package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/internal/dto"
)

type IAuthService interface {
	Register(ctx context.Context, request *dto.RegisterRequest) (dto.RegisterResponse, error)
}
type IAuthHandler interface {
	Register(*gin.Context)
}
