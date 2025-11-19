package file

import (
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
)

type fileUsecase struct {
	fileRepo FileRepository
}

func NewFileUsecase(fileRepo FileRepository) FileUsecase {
	return &fileUsecase{fileRepo}
}

func (u *fileUsecase) UploadFile(c *gin.Context, file *multipart.FileHeader, module string, ownerId uuid.UUID) (*domain.File, error) {
	// Buat direktori files jika belum ada
	module = strings.ReplaceAll(module, " ", "-")
	module = strings.ToLower(module)

	err := u.fileRepo.DeleteFileByModule(module)
	if err != nil {
		log.Println("Error deleting file:", err)
	}

	os.RemoveAll("files/" + module)

	fileDir := "files/" + module
	if err := os.MkdirAll(fileDir, 0755); err != nil {
		return nil, err
	}

	// Generate nama file unik
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	filepath := filepath.Join(fileDir, filename)

	// Buat record file
	fileRecord := &domain.File{
		Name:    filename,
		Type:    file.Header.Get("Content-Type"),
		Size:    file.Size,
		Path:    filepath,
		Module:  module,
		OwnerId: ownerId,
	}

	// Simpan ke database
	if err := u.fileRepo.CreateFile(fileRecord); err != nil {
		return nil, err
	}

	// Simpan file fisik
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		return nil, err
	}

	return fileRecord, nil
}

func (u *fileUsecase) GetFileByID(id uuid.UUID) (*domain.File, error) {
	return u.fileRepo.GetFileByID(id)
}

func (u *fileUsecase) DeleteFile(id uuid.UUID) error {
	// Ambil info file sebelum dihapus
	file, err := u.fileRepo.GetFileByID(id)
	if err != nil {
		return err
	}

	// Hapus file fisik
	if err := os.Remove(file.Path); err != nil {
		return err
	}

	// Hapus record dari database
	return u.fileRepo.DeleteFile(id)
}
