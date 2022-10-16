package openController

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers"
	"net/http"
)

func AppsMyApps(c *gin.Context) {
	number := c.Query("number")
	if number == "" {
		c.JSON(http.StatusOK, controllers.Fail("缺少参数number"))
		return
	}
	c.JSON(http.StatusOK, controllers.Succ("hello"))
	return
}
