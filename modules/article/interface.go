package article

import (
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
)

// ArticleUsecase adalah interface untuk logika bisnis article
type ArticleUsecase interface {
	GetArticles(page, limit int, search, category, tag, status, sort, order string) ([]domain.Article, error)
	GetArticleByID(id uuid.UUID) (*domain.Article, error)
	CreateArticle(article *domain.Article) error
	UpdateArticle(article *domain.Article) error
	DeleteArticle(id uuid.UUID) error
	GetArticleBySlug(slug string) (*domain.Article, error)
}

// ArticleRepository adalah interface untuk operasi database article
type ArticleRepository interface {
	GetArticles(page, limit int, search, category, tag, status, sort, order string) ([]domain.Article, error)
	GetArticleByID(id uuid.UUID) (*domain.Article, error)
	CreateArticle(article *domain.Article) error
	UpdateArticle(article *domain.Article) error
	DeleteArticle(id uuid.UUID) error
	GetArticleBySlug(slug string) (*domain.Article, error)
}
