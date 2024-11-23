package auth

import (
	"context"
	"time"

	"github.com/hilmiikhsan/library-auth-service/constants"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/hilmiikhsan/library-auth-service/internal/dto"
	"github.com/hilmiikhsan/library-auth-service/internal/interfaces"
	"github.com/hilmiikhsan/library-auth-service/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	UserRepo        interfaces.IUserRepository
	UserSessionRepo interfaces.IUserSessionRepository
	Logger          *logrus.Logger
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

func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		res dto.LoginResponse
		now = time.Now()
	)

	userData, err := s.UserRepo.FindUserByUsername(ctx, req.Username)
	if err != nil {
		s.Logger.Error("service::Login - Failed to find user by username : ", err)
		return res, err
	}

	if !helpers.ComparePassword(userData.Password, req.Password) {
		s.Logger.Error("service::Login - Password not match")
		return res, errors.New(constants.ErrUsernameOrPasswordIsIncorrect)
	}

	token, err := helpers.GenerateToken(ctx, userData.ID.String(), userData.Username, constants.TokenTypeAccess, now)
	if err != nil {
		s.Logger.Error("service::Login - Failed to generate token : ", err)
		return res, errors.New(constants.ErrFailedGenerateToken)
	}

	refreshToken, err := helpers.GenerateToken(ctx, userData.ID.String(), userData.Username, constants.RefreshTokenAccess, now)
	if err != nil {
		s.Logger.Error("service::Login - Failed to generate refresh token : ", err)
		return res, errors.New(constants.ErrFailedGenerateRefreshToken)
	}

	userSession := &models.UserSession{
		UserID:              userData.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken[constants.TokenTypeAccess]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken[constants.RefreshTokenAccess]),
	}

	err = s.UserSessionRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		s.Logger.Error("service::Login - Failed to insert new user session : ", err)
		return res, errors.Wrap(err, "failed to insert new user session")
	}

	res.UserID = userData.ID.String()
	res.Username = userData.Username
	res.Token = token
	res.RefreshToken = refreshToken

	return res, nil
}
