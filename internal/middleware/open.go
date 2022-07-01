package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func OpenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Query("access_token"))
		fmt.Println("OpenMiddleware Before")
		c.Next()
		fmt.Println("OpenMiddleware After")
	}
}
