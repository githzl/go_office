package openController

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers"
	"net/http"
)

func TestTest(c *gin.Context) {
	c.JSON(http.StatusOK, controllers.Succ(c.Query("a")))
	return
}
