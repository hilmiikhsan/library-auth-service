package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token":         time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

var jwtSecret = []byte(GetEnv("APP_SECRET", ""))

func GenerateToken(ctx context.Context, userID, username, fullName, role, tokenType string, now time.Time) (string, error) {
	claimToken := ClaimToken{
		UserID:   userID,
		Username: username,
		FullName: fullName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTypeToken[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	resultToken, err := token.SignedString(jwtSecret)
	if err != nil {
		Logger.Error("helpers::GenerateToken - Error while generating token: ", err)
		return resultToken, fmt.Errorf("Error while generating token: %v", err)
	}

	return resultToken, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {
	var (
		claimToken *ClaimToken
		ok         bool
	)

	jwtToken, err := jwt.ParseWithClaims(token, &ClaimToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			Logger.Error("helpers::ValidateToken - Unexpected signing method: ", t.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return jwtSecret, nil
	})
	if err != nil {
		Logger.Error("helpers::ValidateToken - Error while parsing token: ", err)
		return claimToken, fmt.Errorf("Error while parsing token: %v", err)
	}

	if claimToken, ok = jwtToken.Claims.(*ClaimToken); !ok || !jwtToken.Valid {
		Logger.Error("helpers::ValidateToken - Invalid token")
		return claimToken, fmt.Errorf("invalid token")
	}

	return claimToken, nil
}
