package section

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genuicomm/api/db"
)

// Module adalah struktur untuk modul article
type Module struct {
	db          *db.DB
	router      *gin.Engine
	routerGroup *gin.RouterGroup
}

// NewModule membuat instance baru dari Module
func NewModule(db *db.DB, router *gin.Engine, routerGroup *gin.RouterGroup) *Module {
	return &Module{
		db:          db,
		router:      router,
		routerGroup: routerGroup,
	}
}

// Name mengembalikan nama modul
func (m *Module) Name() string {
	return "section"
}

// Register mendaftarkan handler untuk modul article
func (m *Module) Register() {
	// Inisialisasi Repository
	repo := NewRepository(m.db.DB)

	// Inisialisasi Usecase
	usecase := NewUsecase(repo)

	// Inisialisasi dan daftarkan handler
	NewHandler(m.router, m.routerGroup, usecase)
}
