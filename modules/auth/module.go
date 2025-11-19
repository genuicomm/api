package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/genuicomm/api/db"
	"github.com/genuicomm/api/modules/user"
)

// Module adalah struktur untuk modul article
type Module struct {
	db     *db.DB
	router *gin.Engine
}

// NewModule membuat instance baru dari Module
func NewModule(db *db.DB, router *gin.Engine) *Module {
	return &Module{
		db:     db,
		router: router,
	}
}

// Name mengembalikan nama modul
func (m *Module) Name() string {
	return "auth"
}

// Register mendaftarkan handler untuk modul article
func (m *Module) Register() {
	userRepo := user.NewUserRepository(m.db.DB)
	// Inisialisasi Repository

	// Inisialisasi Usecase
	usecase := NewAuthUsecase(userRepo)

	// Inisialisasi dan daftarkan handler
	NewAuthHandler(usecase, m.router)
}
