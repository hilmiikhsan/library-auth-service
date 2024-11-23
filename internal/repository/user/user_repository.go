package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	var res = new(models.User)

	err := r.DB.QueryRowContext(ctx, r.DB.Rebind(queryInsertNewUser),
		user.Username,
		user.Password,
	).Scan(
		&res.ID,
		&res.Username,
	)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if !ok {
			r.Logger.Error("repo::Register - Failed to insert user : ", err)
			return nil, err
		}

		switch pqErr.Code.Name() {
		case "unique_violation":
			r.Logger.Error("repo::Register - Username already registered")
			return nil, errors.New(constants.ErrUsernameAlreadyRegistered)
		default:
			r.Logger.Error("repo::Register - Failed to insert user : ", err)
			return nil, err
		}
	}

	return res, nil
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var res = new(models.User)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindUserByUsername), username)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("repo::FindByUsername - User not found : ", err)
			return nil, errors.New(constants.ErrUsernameOrPasswordIsIncorrect)
		}

		r.Logger.Error("repo::FindByUsername - Failed to find user : ", err)

		return nil, err
	}

	return res, nil
}
