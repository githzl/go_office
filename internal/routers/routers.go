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
		apiGroup.GET("/test/test", apiController.TestTest) // 测试
		apiGroup.GET("/home/get-tags", apiController.HomeGetTags)
		apiGroup.GET("/home/get-tagss", apiController.HomeGetTags)
	}
	return r
}