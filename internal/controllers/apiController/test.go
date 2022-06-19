package apiController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_office/internal/redis"
)

func TestTest(c *gin.Context) {
	var m1 = map[string]string{
		"name": "zhangsan",
		"age":  "18",
	}
	byteStr, _ := json.Marshal(m1)
	str, err := redis.S("demo", byteStr, 10)
	fmt.Println(str, err.Error())
}
