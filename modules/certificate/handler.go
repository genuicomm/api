package certificate

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"github.com/genuicomm/api/pkg/utils"
)

type CertificateHandler struct {
	usecase CertificateUsecase
}

func NewCertificateHandler(router *gin.Engine, routerGroup *gin.RouterGroup, usecase CertificateUsecase) {
	handler := &CertificateHandler{usecase: usecase}

	router.GET("/certificates/:number/validate", handler.ValidateCertificate)

	routerGroup.GET("/certificates", handler.GetAllCertificates)
	routerGroup.GET("/certificates/:id", handler.GetCertificateByID)
	routerGroup.POST("/certificates", handler.CreateCertificate)
	routerGroup.PUT("/certificates/:id", handler.UpdateCertificate)
	routerGroup.GET("/certificates/:id/:status", handler.UpdateCertificateStatus)
	routerGroup.DELETE("/certificates/:id", handler.DeleteCertificate)
}

// GET: Ambil semua certificate
func (h *CertificateHandler) GetAllCertificates(c *gin.Context) {
	certs, err := h.usecase.GetAllCertificates(c.Request.Context())
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, "Daftar sertifikat berhasil diambil", gin.H{"certificates": certs})
}

// GET: Ambil certificate berdasarkan ID
func (h *CertificateHandler) GetCertificateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	cert, err := h.usecase.GetCertificateByID(c.Request.Context(), id.String())
	if err != nil {
		utils.NotFound(c, "Sertifikat tidak ditemukan")
		return
	}

	utils.Success(c, "Sertifikat berhasil diambil", gin.H{"certificate": cert})
}

// POST: Tambah certificate baru
func (h *CertificateHandler) CreateCertificate(c *gin.Context) {
	var cert domain.Certificate
	if err := c.ShouldBindJSON(&cert); err != nil {
		utils.ValidationError(c, "Data tidak valid", err.Error())
		return
	}

	cert.ID = uuid.New()
	err := h.usecase.CreateCertificate(c.Request.Context(), &cert)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}

	utils.Created(c, "Sertifikat berhasil dibuat", gin.H{"certificate": cert})
}

// PUT: Update certificate
func (h *CertificateHandler) UpdateCertificate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	var cert domain.Certificate
	if err := c.ShouldBindJSON(&cert); err != nil {
		utils.ValidationError(c, "Data tidak valid", err.Error())
		return
	}

	cert.ID = id
	err = h.usecase.UpdateCertificate(c.Request.Context(), &cert)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Sertifikat berhasil diperbarui", gin.H{"certificate": cert})
}

// DELETE: Hapus certificate berdasarkan ID
func (h *CertificateHandler) DeleteCertificate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	err = h.usecase.DeleteCertificate(c.Request.Context(), id.String())
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Sertifikat berhasil dihapus", nil)
}

// PUT: Update status certificate
func (h *CertificateHandler) UpdateCertificateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.BadRequest(c, "ID tidak valid")
		return
	}

	status := c.Param("status")
	if status != string(domain.CertificateStatusIssued) && status != string(domain.CertificateStatusCancelled) {
		utils.BadRequest(c, "Status tidak valid")
		return
	}

	err = h.usecase.UpdateCertificateStatus(c.Request.Context(), id.String(), status)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Status sertifikat berhasil diperbarui", nil)

}

// GET: Validasi sertifikat berdasarkan nomor sertifikat
func (h *CertificateHandler) ValidateCertificate(c *gin.Context) {
	number := c.Param("number")
	valid, err := h.usecase.ValidateCertificate(c.Request.Context(), number)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Sertifikat berhasil divalidasi", gin.H{"certificate": valid})
}
