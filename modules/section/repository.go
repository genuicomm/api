package section

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, section *domain.Section) error {
	now := time.Now()
	section.CreatedAt = now
	section.UpdatedAt = now

	return r.db.WithContext(ctx).Create(section).Error
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Section, error) {
	var section domain.Section
	err := r.db.WithContext(ctx).First(&section, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &section, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.Section, error) {
	var sections []*domain.Section
	err := r.db.Debug().WithContext(ctx).Order("sections.order ASC").Find(&sections).Error
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (r *repository) Update(ctx context.Context, section *domain.Section) error {
	section.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(section).Error
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Section{}, id).Error
}
