package routers

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers/apiController"
	"go_office/internal/controllers/openController"
	"go_office/internal/middleware"
)

func Load() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api").Use(middleware.ApiMiddleware())
	{
		apiGroup.GET("/test/test", apiController.TestTest) // 测试
		apiGroup.GET("/home/get-tags", apiController.HomeGetTags)
	}
	openGroup := r.Group("/open").Use(middleware.OpenMiddleware())
	{
		openGroup.GET("/token/get", openController.TokenGet)      // 获取access_token
		openGroup.GET("/apps/my-apps", openController.AppsMyApps) // 获取我的收藏
	}
	return r
}
