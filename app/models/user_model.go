package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID    `gorm:"type:varchar(36);primary_key;"`
	Name            string       `gorm:"type:varchar(255);not null"`
	Email           string       `gorm:"type:varchar(255);unique;not null"`
	EmailVerifiedAt sql.NullTime `gorm:"type:timestamp null"`
	Password        string       `gorm:"type:varchar(255);not null"`
	RefreshTokens   RefreshToken `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt       sql.NullTime `gorm:"autoCreateTime"`
	UpdatedAt       sql.NullTime `gorm:"autoUpdateTime"`
	DeletedAt       sql.NullTime `gorm:"index"`
}
