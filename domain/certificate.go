package domain

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Number          string    `json:"number" gorm:"type:varchar(100);uniqueIndex"`
	DateOfIssue     time.Time `json:"date_of_issue" gorm:"type:date"`
	ParticipantName string    `json:"participant_name" gorm:"type:varchar(255);not null;index"`
	TrainingTitle   string    `json:"training_title" gorm:"type:varchar(255);not null"`
	Location        string    `json:"location" gorm:"type:varchar(255)"`
	Subject         string    `json:"subject" gorm:"type:text"`
	Signer          string    `json:"signer" gorm:"type:varchar(255)"`
	SignerTitle     string    `json:"signer_title" gorm:"type:varchar(255)"`
	File            string    `json:"file" gorm:"type:varchar(255)"`
	Status          string    `json:"status" gorm:"type:varchar(255)"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type CertificateStatus string

const (
	CertificateStatusIssued    CertificateStatus = "issued"
	CertificateStatusCancelled CertificateStatus = "cancelled"
)
