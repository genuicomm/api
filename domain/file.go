package domain

import (
	"time"

	"github.com/google/uuid"
)

// Upload adalah struktur data untuk file yang diupload
type File struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Type      string    `json:"type" gorm:"type:varchar(100);not null"`
	Size      int64     `json:"size" gorm:"not null"`
	Path      string    `json:"path" gorm:"type:varchar(255);not null"`
	Module    string    `json:"module" gorm:"type:varchar(255);not null"`
	OwnerId   uuid.UUID `json:"owner_id" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
