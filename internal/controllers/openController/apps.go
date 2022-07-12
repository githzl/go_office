package openController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppsMyApps(c *gin.Context) {
	number := c.Query("number")
	if number == "" {
		c.JSON(http.StatusOK, Response.Fail("缺少参数number"))
		return
	}
	c.JSON(http.StatusOK, Response.Succ("hello"))
	return
}
