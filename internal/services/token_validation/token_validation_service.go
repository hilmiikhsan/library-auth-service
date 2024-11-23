package token_validation

import (
	"context"

	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type TokenValidationService struct {
	UserSessionRepo interfaces.IUserSessionRepository
	Logger          *logrus.Logger
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
		err        error
	)

	claimToken, err = helpers.ValidateToken(ctx, token)
	if err != nil {
		s.Logger.Error("service::TokenValidation - Failed to validate token : ", err)
		return claimToken, errors.Wrap(err, "failed to validate token")
	}

	_, err = s.UserSessionRepo.FindUserSessionByToken(ctx, token)
	if err != nil {
		s.Logger.Error("service::TokenValidation - Failed to find user session : ", err)
		return claimToken, errors.Wrap(err, "failed to find user session")
	}

	return claimToken, nil
}
