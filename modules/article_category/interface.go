package articlecategory

import (
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
)

type ArticleCategoryUsecase interface {
	GetCategories(page, limit int, search, sort, order string) ([]domain.ArticleCategory, error)
	GetCategoryByID(id uuid.UUID) (*domain.ArticleCategory, error)
	GetCategoryBySlug(slug string) (*domain.ArticleCategory, error)
	CreateCategory(category *domain.ArticleCategory) error
	UpdateCategory(category *domain.ArticleCategory) error
	DeleteCategory(id uuid.UUID) error
}

type ArticleCategoryRepository interface {
	GetCategories(page, limit int, search, sort, order string) ([]domain.ArticleCategory, error)
	GetCategoryByID(id uuid.UUID) (*domain.ArticleCategory, error)
	GetCategoryBySlug(slug string) (*domain.ArticleCategory, error)
	CreateCategory(category *domain.ArticleCategory) error
	UpdateCategory(category *domain.ArticleCategory) error
	DeleteCategory(id uuid.UUID) error
}
