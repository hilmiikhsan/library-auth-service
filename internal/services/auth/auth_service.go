package auth

import (
	"context"

	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/dto"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	"github.com/hilmiikhsan/library-auth-service/internal/models"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	UserRepo interfaces.IUserRepository
	Logger   *logrus.Logger
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (dto.RegisterResponse, error) {
	var res dto.RegisterResponse

	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		s.Logger.Error("service::Register - Failed to hash password : ", err)
		return res, err
	}
	req.Password = hashPassword

	user, err := s.UserRepo.InsertNewUser(ctx, &models.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		s.Logger.Error("service::Register - Failed to insert new user : ", err)
		return res, err
	}

	res.ID = user.ID.String()
	res.Username = user.Username

	return res, nil
}
