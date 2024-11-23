package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/dto"
)

type IAuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (dto.LoginResponse, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (dto.RefreshTokenResponse, error)
}
type IAuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
	RefreshToken(*gin.Context)
}
