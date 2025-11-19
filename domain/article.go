package domain

import (
	"time"

	"github.com/google/uuid"
)

// Article adalah struktur data untuk artikel
type Article struct {
	ID          uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string           `json:"title" gorm:"type:varchar(255);not null"`
	Slug        string           `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null"`
	Content     string           `json:"content" gorm:"type:text;not null"`
	CategoryID  uuid.UUID        `json:"category_id" gorm:"type:uuid"`
	Category    *ArticleCategory `json:"category" gorm:"foreignKey:CategoryID"`
	Tags        []string         `json:"tags" gorm:"type:text;serializer:json"`
	Status      string           `json:"status" gorm:"type:varchar(50);not null;default:'draft'"`
	CreatedAt   time.Time        `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time        `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	PublishedAt time.Time        `json:"published_at"`
}
