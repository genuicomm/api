package articlecategory

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genuicomm/api/db"
)

type Module struct {
	db          *db.DB
	routerGroup *gin.RouterGroup
}

func NewModule(db *db.DB, routerGroup *gin.RouterGroup) *Module {
	return &Module{
		db:          db,
		routerGroup: routerGroup,
	}
}

func (m *Module) Name() string {
	return "article_category"
}

func (m *Module) Register() {
	repo := NewArticleCategoryRepository(m.db.DB)
	usecase := NewArticleCategoryUsecase(repo)
	NewArticleCategoryHandler(m.routerGroup, usecase)
}
