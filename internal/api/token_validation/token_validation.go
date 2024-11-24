package token_validation

import (
	"context"
	"fmt"

	"github.com/hilmiikhsan/library-auth-service/cmd/proto/tokenvalidation"
	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	tokenvalidation.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.GetToken()
	)

	if token == "" {
		err := fmt.Errorf("handler::ValidateToken - Token is empty")
		helpers.Logger.Error(err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claimsToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		helpers.Logger.Error("handler::ValidateToken - Failed to validate token : ", err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &tokenvalidation.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &tokenvalidation.UserData{
			UserId:   claimsToken.UserID,
			Username: claimsToken.Username,
			FullName: claimsToken.FullName,
			Role:     claimsToken.Role,
		},
	}, nil
}
