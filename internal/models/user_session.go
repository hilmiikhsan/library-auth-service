package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	ID                  uuid.UUID `db:"id"`
	UserID              uuid.UUID `db:"user_id"`
	Token               string    `db:"token"`
	RefreshToken        string    `db:"refresh_token"`
	TokenExpired        time.Time `db:"token_expired"`
	RefreshTokenExpired time.Time `db:"refresh_token_expired"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}
