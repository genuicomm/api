package section

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
)

type Repository interface {
	Create(ctx context.Context, section *domain.Section) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Section, error)
	GetAll(ctx context.Context) ([]*domain.Section, error)
	Update(ctx context.Context, section *domain.Section) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Usecase interface {
	Create(ctx context.Context, section *domain.Section) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Section, error)
	GetAll(ctx context.Context) ([]*domain.Section, error)
	Update(ctx context.Context, section *domain.Section) error
	Delete(ctx context.Context, id uuid.UUID) error
}
