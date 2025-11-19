package article

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
)

type articleUsecase struct {
	repo ArticleRepository
}

// NewArticleUsecase membuat instance baru dari ArticleUsecase
func NewArticleUsecase(repo ArticleRepository) ArticleUsecase {
	return &articleUsecase{repo: repo}
}

// GetArticles mengambil daftar artikel dengan filter dan pagination
func (u *articleUsecase) GetArticles(page, limit int, search, category, tag, status, sort, order string) ([]domain.Article, error) {
	// TODO: Implementasi validasi dan transformasi data
	return u.repo.GetArticles(page, limit, search, category, tag, status, sort, order)
}

// GetArticleByID mengambil artikel berdasarkan ID
func (u *articleUsecase) GetArticleByID(id uuid.UUID) (*domain.Article, error) {
	// TODO: Implementasi validasi ID
	return u.repo.GetArticleByID(id)
}

// CreateArticle membuat artikel baru
func (u *articleUsecase) CreateArticle(article *domain.Article) error {
	// TODO: Implementasi validasi dan transformasi data
	return u.repo.CreateArticle(article)
}

// UpdateArticle memperbarui artikel yang ada
func (u *articleUsecase) UpdateArticle(article *domain.Article) error {
	// TODO: Implementasi validasi dan transformasi data
	return u.repo.UpdateArticle(article)
}

// DeleteArticle menghapus artikel
func (u *articleUsecase) DeleteArticle(id uuid.UUID) error {
	// TODO: Implementasi validasi ID
	return u.repo.DeleteArticle(id)
}

// GetArticleBySlug mengambil artikel berdasarkan slug
func (u *articleUsecase) GetArticleBySlug(slug string) (*domain.Article, error) {
	if slug == "" {
		return nil, errors.New("slug cannot be empty")
	}
	return u.repo.GetArticleBySlug(slug)
}
