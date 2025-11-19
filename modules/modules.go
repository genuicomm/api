package modules

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/genuicomm/api/db"
	"github.com/genuicomm/api/modules/article"
	articlecategory "github.com/genuicomm/api/modules/article_category"
	"github.com/genuicomm/api/modules/auth"
	"github.com/genuicomm/api/modules/certificate"
	f "github.com/genuicomm/api/modules/file"
	"github.com/genuicomm/api/modules/section"
	"github.com/genuicomm/api/modules/user"
	"github.com/genuicomm/api/pkg/config"
	"github.com/genuicomm/api/pkg/middleware"
)

// Module adalah interface yang harus diimplementasikan oleh setiap modul
type Module interface {
	Register()
	Name() string
}

// ModuleManager mengelola semua modul dalam aplikasi
type ModuleManager struct {
	db          *db.DB
	config      *config.Config
	modules     map[string]Module
	router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

// NewModuleManager membuat instance baru ModuleManager
func NewModuleManager(cfg *config.Config, db *db.DB) *ModuleManager {
	router := gin.Default()

	// Daftarkan middleware global
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	routerGroup := router.Group("/v1")
	routerGroup.Use(middleware.AuthMiddleware())

	return &ModuleManager{
		db:          db,
		config:      cfg,
		modules:     make(map[string]Module),
		router:      router,
		RouterGroup: routerGroup,
	}
	// Route untuk akses file statis
}

// RegisterModule mendaftarkan modul baru ke ModuleManager
func (mm *ModuleManager) RegisterModule(module Module) error {
	name := module.Name()
	if _, exists := mm.modules[name]; exists {
		return fmt.Errorf("modul %s sudah terdaftar", name)
	}

	mm.modules[name] = module
	log.Printf("Modul %s berhasil didaftarkan", name)
	return nil
}

// GetModule mengembalikan modul berdasarkan nama
func (mm *ModuleManager) GetModule(name string) (Module, error) {
	module, exists := mm.modules[name]
	if !exists {
		return nil, fmt.Errorf("modul %s tidak ditemukan", name)
	}
	return module, nil
}

// InitializeModules menginisialisasi semua modul yang terdaftar
func (mm *ModuleManager) InitializeModules() {
	for name, module := range mm.modules {
		log.Printf("======Menginisialisasi modul: %s=======", name)
		module.Register()
	}
}

// GetRouter mengembalikan router yang sudah dikonfigurasi
func (mm *ModuleManager) GetRouter() *gin.Engine {
	return mm.router
}

func (mm *ModuleManager) GetRouterGroup() *gin.RouterGroup {
	return mm.RouterGroup
}

// InitModules menginisialisasi semua modul dan mengembalikan router yang sudah dikonfigurasi
func InitModules(cfg *config.Config, db *db.DB) *gin.Engine {
	// Buat ModuleManager
	moduleManager := NewModuleManager(cfg, db)

	// Daftarkan modul-modul
	modules := []Module{
		auth.NewModule(db, moduleManager.GetRouter()),
		user.NewModule(db, moduleManager.GetRouterGroup()),
		article.NewModule(db, moduleManager.GetRouter(), moduleManager.GetRouterGroup()),
		articlecategory.NewModule(db, moduleManager.GetRouterGroup()),
		certificate.NewModule(db, moduleManager.GetRouter(), moduleManager.GetRouterGroup()),
		section.NewModule(db, moduleManager.GetRouter(), moduleManager.GetRouterGroup()),
		f.NewModule(db, moduleManager.GetRouter(), moduleManager.GetRouterGroup()),
		// Tambahkan modul baru di sini
	}

	// Register semua modul
	for _, module := range modules {
		if err := moduleManager.RegisterModule(module); err != nil {
			log.Printf("Gagal mendaftarkan modul %s: %v", module.Name(), err)
			continue
		}
	}

	// Inisialisasi semua modul
	moduleManager.InitializeModules()

	return moduleManager.GetRouter()
}
