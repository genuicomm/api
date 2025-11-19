package section

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"github.com/genuicomm/api/pkg/utils"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(router *gin.Engine, routerGroup *gin.RouterGroup, usecase Usecase) {
	handler := &Handler{usecase: usecase}

	router.GET("/sections", handler.GetAll)

	routerGroup.GET("/sections", handler.GetAll)
	routerGroup.GET("/sections/:id", handler.GetByID)
	routerGroup.POST("/sections", handler.Create)
	routerGroup.PUT("/sections/:id", handler.Update)
	routerGroup.DELETE("/sections/:id", handler.Delete)
}

func (h *Handler) Create(c *gin.Context) {
	var section domain.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		utils.ValidationError(c, "Data tidak valid", err.Error())
		return
	}

	if err := h.usecase.Create(c.Request.Context(), &section); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Created(c, "Section berhasil dibuat", gin.H{"section": section})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	section, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	if section == nil {
		utils.NotFound(c, "Section tidak ditemukan")
		return
	}

	utils.Success(c, "Section berhasil diambil", gin.H{"section": section})
}

func (h *Handler) GetAll(c *gin.Context) {
	sections, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Daftar section berhasil diambil", gin.H{"sections": sections})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	var section domain.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		utils.ValidationError(c, "Data tidak valid", err.Error())
		return
	}

	section.ID = id
	if err := h.usecase.Update(c.Request.Context(), &section); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Section berhasil diperbarui", gin.H{"section": section})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	if err := h.usecase.Delete(c.Request.Context(), id); err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Success(c, "Section berhasil dihapus", nil)
}
