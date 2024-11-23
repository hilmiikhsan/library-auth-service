package interfaces

import (
	"context"

	"github.com/hilmiikhsan/library-auth-service/cmd/proto/tokenvalidation"
	"github.com/hilmiikhsan/library-auth-service/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
