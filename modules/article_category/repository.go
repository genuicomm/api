package articlecategory

import (
	"strings"

	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"gorm.io/gorm"
)

type articleCategoryRepository struct {
	db *gorm.DB
}

func NewArticleCategoryRepository(db *gorm.DB) ArticleCategoryRepository {
	return &articleCategoryRepository{db: db}
}

func (r *articleCategoryRepository) GetCategories(page, limit int, search, sort, order string) ([]domain.ArticleCategory, error) {
	var categories []domain.ArticleCategory
	query := r.db.Model(&domain.ArticleCategory{})

	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort + " " + order)
	} else {
		query = query.Order("created_at DESC")
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&categories).Error

	return categories, err
}

func (r *articleCategoryRepository) GetCategoryByID(id uuid.UUID) (*domain.ArticleCategory, error) {
	var category domain.ArticleCategory
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *articleCategoryRepository) GetCategoryBySlug(slug string) (*domain.ArticleCategory, error) {
	var category domain.ArticleCategory
	err := r.db.First(&category, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *articleCategoryRepository) CreateCategory(category *domain.ArticleCategory) error {
	if category.Slug == "" {
		category.Slug = generateSlug(category.Name)
	}
	return r.db.Create(category).Error
}

func (r *articleCategoryRepository) UpdateCategory(category *domain.ArticleCategory) error {
	if category.Slug == "" {
		category.Slug = generateSlug(category.Name)
	}
	return r.db.Save(category).Error
}

func (r *articleCategoryRepository) DeleteCategory(id uuid.UUID) error {
	return r.db.Delete(&domain.ArticleCategory{}, "id = ?", id).Error
}

func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
