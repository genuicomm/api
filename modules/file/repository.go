package file

import (
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"gorm.io/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db}
}

func (r *fileRepository) CreateFile(file *domain.File) error {
	return r.db.Create(file).Error
}

func (r *fileRepository) GetFileByID(id uuid.UUID) (*domain.File, error) {
	var file domain.File
	err := r.db.First(&file, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) DeleteFile(id uuid.UUID) error {
	return r.db.Delete(&domain.File{}, "id = ?", id).Error
}

func (r *fileRepository) DeleteFileByModule(module string) error {
	return r.db.Delete(&domain.File{}, "module = ?", module).Error
}
