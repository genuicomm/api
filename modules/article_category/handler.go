package articlecategory

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
	"gitlab.com/genuicomm/api/pkg/utils"
	"gorm.io/gorm"
)

type ArticleCategoryHandler struct {
	usecase ArticleCategoryUsecase
}

func NewArticleCategoryHandler(routerGroup *gin.RouterGroup, usecase ArticleCategoryUsecase) {
	handler := &ArticleCategoryHandler{
		usecase: usecase,
	}

	routerGroup.GET("/article-categories", handler.GetCategories)
	routerGroup.GET("/article-categories/:id", handler.GetCategoryByID)
	routerGroup.GET("/article-categories/slug/:slug", handler.GetCategoryBySlug)
	routerGroup.POST("/article-categories", handler.CreateCategory)
	routerGroup.PUT("/article-categories/:id", handler.UpdateCategory)
	routerGroup.DELETE("/article-categories/:id", handler.DeleteCategory)
}

func (h *ArticleCategoryHandler) GetCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	categories, err := h.usecase.GetCategories(page, limit, search, sort, order)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Daftar kategori berhasil diambil", gin.H{"categories": categories})
}

func (h *ArticleCategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "ID kategori tidak valid")
		return
	}

	category, err := h.usecase.GetCategoryByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Kategori tidak ditemukan")
			return
		}
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Kategori berhasil diambil", gin.H{"category": category})
}

func (h *ArticleCategoryHandler) GetCategoryBySlug(c *gin.Context) {
	slug := c.Param("slug")

	category, err := h.usecase.GetCategoryBySlug(slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Kategori tidak ditemukan")
			return
		}
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Kategori berhasil diambil", gin.H{"category": category})
}

func (h *ArticleCategoryHandler) CreateCategory(c *gin.Context) {
	var category domain.ArticleCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.CreateCategory(&category); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Kategori berhasil dibuat", gin.H{"category": category})
}

func (h *ArticleCategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "ID kategori tidak valid")
		return
	}

	var category domain.ArticleCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	category.ID = id
	if err := h.usecase.UpdateCategory(&category); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Kategori berhasil diperbarui", gin.H{"category": category})
}

func (h *ArticleCategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "ID kategori tidak valid")
		return
	}

	if err := h.usecase.DeleteCategory(id); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Kategori berhasil dihapus", nil)
}
