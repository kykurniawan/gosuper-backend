package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Otp struct {
	ID         uuid.UUID    `gorm:"primaryKey"`
	Code       string       `gorm:"unique"`
	OtpType    string       `gorm:"not null"`
	UserID     uuid.UUID    `gorm:"not null"`
	ValidUntil time.Time    `gorm:"not null"`
	CreatedAt  sql.NullTime `gorm:"autoCreateTime"`
	UpdatedAt  sql.NullTime `gorm:"autoUpdateTime"`
}
