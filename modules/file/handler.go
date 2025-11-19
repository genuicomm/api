package file

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileHandler struct {
	usecase FileUsecase
}

func NewFileHandler(router *gin.Engine, routerGroup *gin.RouterGroup, usecase FileUsecase) {
	handler := &FileHandler{usecase: usecase}

	fileGroup := routerGroup.Group("/files")
	{
		fileGroup.POST("", handler.UploadFile)
		fileGroup.GET("/:id", handler.GetFile)
		fileGroup.DELETE("/:id", handler.DeleteFile)
	}

	router.GET("/files/:category/:filename", handler.ServeFile)
}

// UploadFile menangani request untuk upload file
func (h *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	module := c.PostForm("module")
	ownerId, err := uuid.Parse(c.PostForm("owner_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Upload file menggunakan usecase
	fileRecord, err := h.usecase.UploadFile(c, file, module, ownerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	c.JSON(http.StatusOK, fileRecord)
}

// GetFile menangani request untuk mendapatkan file berdasarkan ID
func (h *FileHandler) GetFile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	file, err := h.usecase.GetFileByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, file)
}

// DeleteFile menangani request untuk menghapus file
func (h *FileHandler) DeleteFile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := h.usecase.DeleteFile(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File berhasil dihapus"})
}

// ServeFile menangani request untuk mengakses file berdasarkan kategori dan nama file
func (h *FileHandler) ServeFile(c *gin.Context) {
	category := c.Param("category")
	filename := c.Param("filename")

	filePath := filepath.Join("files", category, filename)

	// Cek apakah file ada
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File tidak ditemukan"})
		return
	}

	// Serve file
	c.File(filePath)
}
