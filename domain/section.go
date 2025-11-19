package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Section struct {
	ID        uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string          `json:"name" gorm:"type:varchar(255);not null"`
	Order     int             `json:"order" gorm:"type:integer;not null;default:0"`
	Type      string          `json:"type" gorm:"type:varchar(255);not null"`
	Data      json.RawMessage `json:"data" gorm:"type:jsonb;not null"`
	IsActive  bool            `json:"is_active" gorm:"type:boolean;not null;default:true"`
	CreatedAt time.Time       `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type SectionData struct {
	IsVisible     bool   `json:"is_visible"`
	Label         string `json:"label"`
	Key           string `json:"key"`
	Type          string `json:"type"`
	Class         string `json:"class"`
	Val           any    `json:"value"`
	Description   string `json:"description"`
	Image         string `json:"image"`
	LabelEditable bool   `json:"label_editable"`
}

type SectionDataList []SectionData

// Scan implements the sql.Scanner interface
func (s *SectionDataList) Scan(value interface{}) error {
	if value == nil {
		*s = make(SectionDataList, 0)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(bytes, s)
}

// Value implements the driver.Valuer interface
func (s SectionDataList) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}

	return json.Marshal(s)
}
