package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Token      string    `gorm:"unique"`
	ValidUntil time.Time `gorm:"not null"`
	UserID     uuid.UUID
	CreatedAt  sql.NullTime `gorm:"autoCreateTime"`
	UpdatedAt  sql.NullTime `gorm:"autoUpdateTime"`
}
