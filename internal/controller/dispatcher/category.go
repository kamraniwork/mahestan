package dispatcher

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mahestan/internal/controller"
	"mahestan/internal/repository"
	"mahestan/internal/service"
)

func SetupCategoriesRoutes(gormDb *gorm.DB) *controller.CategoryController {
	categoryRepository := repository.NewCategoriesRepository(gormDb)
	categorySvc := service.NewCategoryService(
		categoryRepository,
	)
	return controller.NewCategory(
		*categorySvc,
	)
}

func SetupCategoryRoutes(routes gin.IRoutes, categoryCtrl *controller.CategoryController) {
	routes.POST("/category/create", categoryCtrl.CreateCategory)
}
