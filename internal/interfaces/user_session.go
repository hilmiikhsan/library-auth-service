package interfaces

import (
	"context"

	"github.com/hilmiikhsan/library-auth-service/internal/models"
)

type IUserSessionRepository interface {
	FindUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
}
