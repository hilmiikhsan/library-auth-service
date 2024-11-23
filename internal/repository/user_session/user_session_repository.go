package user_session

import (
	"context"

	"github.com/hilmiikhsan/library-auth-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserSessionRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r UserSessionRepository) FindUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var res = new(models.UserSession)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindUserSessionByToken), token)
	if err != nil {
		r.Logger.Error("repo::FindUserSessionByToken - Failed to find user session by token : ", err)
		return nil, err
	}

	return res, nil
}

func (r UserSessionRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryInsertNewUserSession),
		session.UserID,
		session.Token,
		session.RefreshToken,
		session.TokenExpired,
		session.RefreshTokenExpired,
	)
	if err != nil {
		r.Logger.Error("repo::InsertNewUserSession - Failed to insert new user session : ", err)
		return err
	}

	return nil
}
