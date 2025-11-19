package article

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genuicomm/api/db"
)

// Module adalah struktur untuk modul article
type Module struct {
	db          *db.DB
	router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

// NewModule membuat instance baru dari Module
func NewModule(db *db.DB, router *gin.Engine, group *gin.RouterGroup) *Module {
	return &Module{
		db:          db,
		router:      router,
		RouterGroup: group,
	}
}

// Name mengembalikan nama modul
func (m *Module) Name() string {
	return "article"
}

// Register mendaftarkan handler untuk modul article
func (m *Module) Register() {
	// Inisialisasi Repository
	repo := NewArticleRepository(m.db.DB)

	// Inisialisasi Usecase
	usecase := NewArticleUsecase(repo)

	// Inisialisasi dan daftarkan handler
	NewArticleHandler(m.router, m.RouterGroup, usecase)
}
