package models

import "github.com/google/uuid"

type UserSession struct {
	ID     uuid.UUID `db:"id"`
	UserID uuid.UUID `db:"user_id"`
}
