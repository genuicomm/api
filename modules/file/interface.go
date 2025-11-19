package file

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
)

type FileUsecase interface {
	UploadFile(c *gin.Context, file *multipart.FileHeader, module string, ownerId uuid.UUID) (*domain.File, error)
	GetFileByID(id uuid.UUID) (*domain.File, error)
	DeleteFile(id uuid.UUID) error
}

type FileRepository interface {
	CreateFile(file *domain.File) error
	GetFileByID(id uuid.UUID) (*domain.File, error)
	DeleteFile(id uuid.UUID) error
	DeleteFileByModule(module string) error
}
