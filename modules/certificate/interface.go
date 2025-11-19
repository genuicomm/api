package certificate

import (
	"context"

	"gitlab.com/genuicomm/api/domain"
)

type CertificateUsecase interface {
	GetAllCertificates(ctx context.Context) ([]domain.Certificate, error)
	GetCertificateByID(ctx context.Context, id string) (domain.Certificate, error)
	CreateCertificate(ctx context.Context, cert *domain.Certificate) error
	UpdateCertificate(ctx context.Context, cert *domain.Certificate) error
	UpdateCertificateStatus(ctx context.Context, id string, status string) error
	DeleteCertificate(ctx context.Context, id string) error
	ValidateCertificate(ctx context.Context, number string) (domain.Certificate, error)
}

type CertificateRepository interface {
	GetAll(ctx context.Context) ([]domain.Certificate, error)
	GetByID(ctx context.Context, id string) (domain.Certificate, error)
	GetByNumber(ctx context.Context, number string) (domain.Certificate, error)
	Create(ctx context.Context, cert *domain.Certificate) error
	Update(ctx context.Context, cert *domain.Certificate) error
	Delete(ctx context.Context, id string) error
}
