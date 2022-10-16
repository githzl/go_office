package apiController

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers"
	"net/http"
)

func TestTest(c *gin.Context) {
	//var m1 = map[string]string{
	//	"name": "zhangsan",
	//	"age":  "18",
	//}
	m1 := c.Query("name")
	c.JSON(http.StatusOK, controllers.Succ(m1))
}
