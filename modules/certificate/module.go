package certificate

import (
	"github.com/gin-gonic/gin"
	"github.com/genuicomm/api/db"
)

type Module struct {
	db          *db.DB
	router      *gin.Engine
	routerGroup *gin.RouterGroup
}

func NewModule(db *db.DB, router *gin.Engine, routerGroup *gin.RouterGroup) *Module {
	return &Module{
		db:          db,
		router:      router,
		routerGroup: routerGroup,
	}
}

func (m *Module) Name() string {
	return "certificate"
}

func (m *Module) Register() {
	// Inisialisasi Repository
	repo := NewCertificateRepository(m.db.DB)

	// Inisialisasi Usecase
	usecase := NewCertificateUsecase(repo)

	// Inisialisasi dan daftarkan handler
	NewCertificateHandler(m.router, m.routerGroup, usecase)
}
