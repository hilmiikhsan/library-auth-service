package interfaces

import (
	"context"

	"github.com/hilmiikhsan/library-auth-service/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) (*models.User, error)
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
}
