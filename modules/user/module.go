package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genuicomm/api/db"
)

// Module adalah struktur untuk modul user
type Module struct {
	db          *db.DB
	routerGroup *gin.RouterGroup
}

// NewModule membuat instance baru dari Module
func NewModule(db *db.DB, routerGroup *gin.RouterGroup) *Module {
	return &Module{
		db:          db,
		routerGroup: routerGroup,
	}
}

// Name mengembalikan nama modul
func (m *Module) Name() string {
	return "user"
}

// Register mendaftarkan handler untuk modul user
func (m *Module) Register() {
	// Inisialisasi Repository
	repo := NewUserRepository(m.db.DB)

	// Inisialisasi Usecase
	usecase := NewUserUsecase(repo)

	// Inisialisasi dan daftarkan handler
	NewUserHandler(usecase, m.routerGroup)
}
