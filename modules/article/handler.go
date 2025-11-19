package article

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"github.com/genuicomm/api/pkg/utils"
	"gorm.io/gorm"
)

// ArticleHandler menangani HTTP request untuk article
type ArticleHandler struct {
	usecase ArticleUsecase
}

// NewArticleHandler membuat instance baru dari ArticleHandler
func NewArticleHandler(router *gin.Engine, routerGroup *gin.RouterGroup, usecase ArticleUsecase) {
	handler := &ArticleHandler{
		usecase: usecase,
	}

	router.GET("/articles", handler.GetArticles)
	router.GET("/articles/:id", handler.GetArticleByID)
	// Daftarkan routes
	routerGroup.GET("/articles", handler.GetArticles)
	routerGroup.GET("/articles/:id", handler.GetArticleByID)
	routerGroup.POST("/articles", handler.CreateArticle)
	routerGroup.PUT("/articles/:id", handler.UpdateArticle)
	routerGroup.DELETE("/articles/:id", handler.DeleteArticle)
	routerGroup.GET("/articles/slug/:slug", handler.GetArticleBySlug)
}

// GetArticles menangani request untuk mengambil daftar artikel
func (h *ArticleHandler) GetArticles(c *gin.Context) {
	// Parse query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")
	category := c.DefaultQuery("category", "")
	tag := c.DefaultQuery("tag", "")
	status := c.DefaultQuery("status", "")
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	articles, err := h.usecase.GetArticles(page, limit, search, category, tag, status, sort, order)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	// Hitung pagination
	total := len(articles) // TODO: Dapatkan total dari database
	lastPage := (total + limit - 1) / limit
	from := (page-1)*limit + 1
	to := from + len(articles) - 1

	pagination := utils.Pagination{
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
		From:        from,
		To:          to,
	}

	response := gin.H{
		"articles":   articles,
		"pagination": pagination,
	}

	utils.Success(c, "Daftar artikel berhasil diambil", response)
}

// GetArticleByID menangani request untuk mengambil artikel berdasarkan ID
func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	article, err := h.usecase.GetArticleByID(uuid.MustParse(id))
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	if article == nil {
		utils.NotFound(c, "Artikel tidak ditemukan")
		return
	}

	utils.Success(c, "Detail artikel berhasil diambil", gin.H{"article": article})
}

// CreateArticle menangani request untuk membuat artikel baru
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article domain.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.ValidationError(c, "Data tidak valid", err.Error())
		return
	}

	if err := h.usecase.CreateArticle(&article); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Created(c, "Artikel berhasil dibuat", gin.H{"article": article})
}

// UpdateArticle menangani request untuk memperbarui artikel
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article domain.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.ValidationError(c, "Data tidak valid", err.Error())
		return
	}

	article.ID = uuid.MustParse(id)
	if err := h.usecase.UpdateArticle(&article); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Artikel berhasil diperbarui", gin.H{"article": article})
}

// DeleteArticle menangani request untuk menghapus artikel
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	if err := h.usecase.DeleteArticle(uuid.MustParse(id)); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Artikel berhasil dihapus", nil)
}

// GetArticleBySlug menangani request untuk mendapatkan artikel berdasarkan slug
func (h *ArticleHandler) GetArticleBySlug(c *gin.Context) {
	slug := c.Param("slug")

	article, err := h.usecase.GetArticleBySlug(slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Artikel tidak ditemukan")
			return
		}
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Detail artikel berhasil diambil", gin.H{"article": article})
}
