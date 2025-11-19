package certificate

import (
	"context"
	"errors"

	"gitlab.com/genuicomm/api/domain"
)

type certificateUsecase struct {
	repo CertificateRepository
}

func NewCertificateUsecase(repo CertificateRepository) CertificateUsecase {
	return &certificateUsecase{repo: repo}
}

// Ambil semua sertifikat
func (u *certificateUsecase) GetAllCertificates(ctx context.Context) ([]domain.Certificate, error) {
	return u.repo.GetAll(ctx)
}

// Ambil sertifikat berdasarkan ID
func (u *certificateUsecase) GetCertificateByID(ctx context.Context, id string) (domain.Certificate, error) {
	return u.repo.GetByID(ctx, id)
}

// Tambah sertifikat baru
func (u *certificateUsecase) CreateCertificate(ctx context.Context, cert *domain.Certificate) error {
	// Validasi jika nama peserta atau judul kosong
	if cert.ParticipantName == "" || cert.TrainingTitle == "" {
		return errors.New("nama peserta dan judul pelatihan wajib diisi")
	}

	cert.Status = string(domain.CertificateStatusIssued)

	return u.repo.Create(ctx, cert)
}

// Update sertifikat
func (u *certificateUsecase) UpdateCertificate(ctx context.Context, cert *domain.Certificate) error {
	_, err := u.repo.GetByID(ctx, cert.ID.String()) // Pastikan data ada
	if err != nil {
		return errors.New("sertifikat tidak ditemukan")
	}

	return u.repo.Update(ctx, cert)
}

// Hapus sertifikat berdasarkan ID
func (u *certificateUsecase) DeleteCertificate(ctx context.Context, id string) error {
	_, err := u.repo.GetByID(ctx, id) // Pastikan data ada
	if err != nil {
		return errors.New("sertifikat tidak ditemukan")
	}

	return u.repo.Delete(ctx, id)
}

// Update status sertifikat
func (u *certificateUsecase) UpdateCertificateStatus(ctx context.Context, id string, status string) error {
	existingCert, err := u.repo.GetByID(ctx, id) // Pastikan data ada
	if err != nil {
		return errors.New("sertifikat tidak ditemukan")
	}

	if existingCert.Status == status {
		return errors.New("status sertifikat sudah sama")
	}

	existingCert.Status = status

	return u.repo.Update(ctx, &existingCert)
}

// Validasi sertifikat berdasarkan nomor sertifikat
func (u *certificateUsecase) ValidateCertificate(ctx context.Context, number string) (domain.Certificate, error) {
	existingCert, err := u.repo.GetByNumber(ctx, number) // Pastikan data ada
	if err != nil {
		return domain.Certificate{}, errors.New("sertifikat tidak ditemukan")
	}
	if existingCert.Status != string(domain.CertificateStatusIssued) {
		return domain.Certificate{}, errors.New("sertifikat tidak validxxxx")
	}

	return existingCert, nil
}
