package domain

import (
	"time"

	"github.com/google/uuid"
)

// User merepresentasikan model user
type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email        string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password     string    `json:"-" gorm:"-"`
	PasswordHash []byte    `json:"-" gorm:"type:bytea;not null"`
	FullName     string    `json:"full_name" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
