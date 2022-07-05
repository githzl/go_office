package openController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppsMyApps(c *gin.Context) {
	c.JSON(http.StatusOK, Response.Succ("hello"))
	return
}