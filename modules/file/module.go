package file

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genuicomm/api/db"
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
	return "file"
}

func (m *Module) Register() {
	// Inisialisasi Repository
	repo := NewFileRepository(m.db.DB)

	// Inisialisasi Usecase
	usecase := NewFileUsecase(repo)

	// Inisialisasi dan daftarkan handler
	NewFileHandler(m.router, m.routerGroup, usecase)
}
