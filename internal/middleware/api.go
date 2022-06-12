package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ApiMiddleware Before")
		c.Next()
		fmt.Println("ApiMiddleware After")
	}
}
