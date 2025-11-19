package section

import (
	"context"

	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
)

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(ctx context.Context, section *domain.Section) error {
	return u.repo.Create(ctx, section)
}

func (u *usecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Section, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *usecase) GetAll(ctx context.Context) ([]*domain.Section, error) {
	return u.repo.GetAll(ctx)
}

func (u *usecase) Update(ctx context.Context, section *domain.Section) error {
	existingSection, err := u.repo.GetByID(ctx, section.ID)
	if err != nil {
		return err
	}

	existingSection.Name = section.Name
	existingSection.Order = section.Order
	existingSection.Type = section.Type
	existingSection.Data = section.Data

	return u.repo.Update(ctx, existingSection)
}

func (u *usecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
