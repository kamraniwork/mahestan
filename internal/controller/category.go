package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mahestan/internal/controller/serializers"
	"mahestan/internal/service"
	"net/http"
)

type CategoryController struct {
	svc service.CategoryService
}

func NewCategory(categorySvc service.CategoryService) *CategoryController {
	return &CategoryController{
		svc: categorySvc,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req serializers.InputCategorySerializer

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logrus.WithError(err).Error("validation Error!")

		return
	}

	category, err := c.svc.Create(ctx.Request.Context(), req)
	if err != nil {
		logrus.WithError(err).Error("create category Error!")
	}
	output := serializers.OutputDetailCategorySerializer{
		Name:        category.Name,
		Description: category.Description,
	}

	ctx.JSON(http.StatusOK, output)

}
