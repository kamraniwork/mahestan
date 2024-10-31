package dispatcher

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(sqlDB *gorm.DB) (*gin.Engine, error) {
	engine := gin.New()

	v1Group := engine.Group("/v1")
	SetupCategoryRoutes(v1Group, SetupCategoriesRoutes(sqlDB))

	return engine, nil
}
