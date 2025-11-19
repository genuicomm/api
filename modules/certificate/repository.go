package certificate

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"gorm.io/gorm"
)

type certificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &certificateRepository{db: db}
}

// GetAll: Ambil semua data sertifikat
func (r *certificateRepository) GetAll(ctx context.Context) ([]domain.Certificate, error) {
	var certificates []domain.Certificate
	err := r.db.Debug().Find(&certificates).Error
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

// GetByID: Ambil Certificate berdasarkan ID
func (r *certificateRepository) GetByID(ctx context.Context, id string) (domain.Certificate, error) {
	var cert domain.Certificate
	err := r.db.Where("id = ?", id).First(&cert).Error
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Certificate{}, ErrCertificateNotFound
		}
		return domain.Certificate{}, err
	}
	return cert, nil
}

var ErrCertificateNotFound = errors.New("Certificate not found")

// Create: Tambahkan Certificate baru
func (r *certificateRepository) Create(ctx context.Context, cert *domain.Certificate) error {
	// Generate UUID jika belum ada
	if cert.ID == uuid.Nil {
		cert.ID = uuid.New()
	}

	err := r.db.Debug().Create(cert).Error
	if err != nil {
		return err
	}

	return nil
}

// Update: Perbarui data sertifikat
func (r *certificateRepository) Update(ctx context.Context, cert *domain.Certificate) error {
	err := r.db.Model(&domain.Certificate{}).Where("id = ?", cert.ID).Updates(cert).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete: Hapus certificate berdasarkan ID
func (r *certificateRepository) Delete(ctx context.Context, id string) error {
	err := r.db.Model(&domain.Certificate{}).Where("id = ?", id).Delete(&domain.Certificate{}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByNumber: Ambil Certificate berdasarkan nomor sertifikat
func (r *certificateRepository) GetByNumber(ctx context.Context, number string) (domain.Certificate, error) {
	var cert domain.Certificate
	err := r.db.Where("number = ?", number).First(&cert).Error
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Certificate{}, ErrCertificateNotFound
		}
		return domain.Certificate{}, err
	}
	return cert, nil
}
