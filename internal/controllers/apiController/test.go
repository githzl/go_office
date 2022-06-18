package apiController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_office/internal/redis"
)

func TestTest(c *gin.Context)  {
	var m1 = map[string]string{
		"name":"zhangsan",
		"age":"18",
	}
	str, err := redis.S("demo",m1,10)
	fmt.Println(str,err.Error())
}