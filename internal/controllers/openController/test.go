package openController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestTest(c *gin.Context) {
	c.JSON(http.StatusOK, Response.Succ(c.Query("a")))
	return
}
