package articlecategory

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
)

type articleCategoryUsecase struct {
	repo ArticleCategoryRepository
}

func NewArticleCategoryUsecase(repo ArticleCategoryRepository) ArticleCategoryUsecase {
	return &articleCategoryUsecase{repo: repo}
}

func (u *articleCategoryUsecase) GetCategories(page, limit int, search, sort, order string) ([]domain.ArticleCategory, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return u.repo.GetCategories(page, limit, search, sort, order)
}

func (u *articleCategoryUsecase) GetCategoryByID(id uuid.UUID) (*domain.ArticleCategory, error) {
	return u.repo.GetCategoryByID(id)
}

func (u *articleCategoryUsecase) GetCategoryBySlug(slug string) (*domain.ArticleCategory, error) {
	if slug == "" {
		return nil, errors.New("slug cannot be empty")
	}
	return u.repo.GetCategoryBySlug(slug)
}

func (u *articleCategoryUsecase) CreateCategory(category *domain.ArticleCategory) error {
	if category.Name == "" {
		return errors.New("name is required")
	}
	return u.repo.CreateCategory(category)
}

func (u *articleCategoryUsecase) UpdateCategory(category *domain.ArticleCategory) error {
	if category.Name == "" {
		return errors.New("name is required")
	}
	return u.repo.UpdateCategory(category)
}

func (u *articleCategoryUsecase) DeleteCategory(id uuid.UUID) error {
	return u.repo.DeleteCategory(id)
}
