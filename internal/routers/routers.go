package routers

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers/apiController"
	"go_office/internal/middleware"
)

func Load() *gin.Engine{
	r := gin.Default()
	apiGroup := r.Group("/api").Use(middleware.ApiMiddleware())
	{
		apiGroup.GET("/home/site-options", apiController.HomeSiteOption)
		apiGroup.GET("/home/get-tags", apiController.HomeGetTags)
	}
	return r
}