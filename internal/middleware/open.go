package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go_office/internal/controllers"
	"net/http"
)

func OpenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() != "/open/token/get" {
			//如果请求path不是"／open/token/get"则需要验证本次携带的access_token是否有效
			accessToken := c.Query("access_token")
			if err := verifyAccessToken(accessToken); err != nil {
				c.JSON(http.StatusOK, controllers.Response.Fail(err.Error()))
				c.Abort()
			}
		}
		c.Next()
	}
}

// 验证access_token是否有效
func verifyAccessToken(accessToken string) error {
	if accessToken == "" {
		return errors.New("缺少参数access_token")
	}
	if accessToken != "abc" {
		return errors.New("access_token验证失败")
	}
	return nil
}
