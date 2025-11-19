package article

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

// NewArticleRepository membuat instance baru dari ArticleRepository
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// GetArticles mengambil daftar artikel dengan filter dan pagination
func (r *articleRepository) GetArticles(page, limit int, search, category, tag, status, sort, order string) ([]domain.Article, error) {
	var articles []domain.Article
	query := r.db.Model(&domain.Article{})

	// Terapkan filter pencarian
	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Filter berdasarkan kategori
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Filter berdasarkan tag (menggunakan JSONB query karena tags adalah array)
	if tag != "" {
		query = query.Where("tags @> ?", []string{tag})
	}

	// Filter berdasarkan status
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Terapkan pengurutan
	if sort != "" {
		query = query.Order(sort + " " + order)
	} else {
		query = query.Order("created_at DESC")
	}

	// Terapkan pagination
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&articles).Error

	return articles, err
}

// GetArticleByID mengambil artikel berdasarkan ID
func (r *articleRepository) GetArticleByID(id uuid.UUID) (*domain.Article, error) {
	var article domain.Article
	err := r.db.First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// CreateArticle membuat artikel baru
func (r *articleRepository) CreateArticle(article *domain.Article) error {
	// Generate slug dari title jika belum ada
	if article.Slug == "" {
		article.Slug = generateSlug(article.Title)
	}

	// Set status default jika belum ada
	if article.Status == "" {
		article.Status = "draft"
	}

	return r.db.Create(article).Error
}

// UpdateArticle memperbarui artikel yang ada
func (r *articleRepository) UpdateArticle(article *domain.Article) error {
	// Update slug jika title berubah
	if article.Slug == "" {
		article.Slug = generateSlug(article.Title)
	}

	return r.db.Save(article).Error
}

// DeleteArticle menghapus artikel
func (r *articleRepository) DeleteArticle(id uuid.UUID) error {
	return r.db.Delete(&domain.Article{}, "id = ?", id).Error
}

// generateSlug adalah fungsi helper untuk menghasilkan slug dari title
func generateSlug(title string) string {
	// Implementasi sederhana, Anda mungkin ingin menggunakan library seperti gosimple/slug
	// untuk implementasi yang lebih robust
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Tambahkan logika pembersihan slug lainnya sesuai kebutuhan
	return slug
}

// GetArticleBySlug mengambil artikel berdasarkan slug
func (r *articleRepository) GetArticleBySlug(slug string) (*domain.Article, error) {
	var article domain.Article
	err := r.db.First(&article, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}
