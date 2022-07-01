package openController

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers"
	"net/http"
)

var Response = new(controllers.ResponseFormat)

func TestTest(c *gin.Context) {
	c.JSON(http.StatusOK, Response.Succ(c.Query("a")))
	return
}
