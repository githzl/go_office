package apiController

import (
	"github.com/gin-gonic/gin"
	"go_office/config"
	"net/http"
)

func HomeSiteOption(c *gin.Context) {
	c.String(http.StatusOK, "Hello yaml.redis.perfix="+config.YamlFile.Redis.Prefix)
}
func HomeGetTags(c *gin.Context) {
	c.String(http.StatusOK, "Hello yaml.redis.perfix="+config.YamlFile.Redis.Prefix)
}