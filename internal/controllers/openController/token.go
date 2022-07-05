package openController

import (
	"github.com/gin-gonic/gin"
	"go_office/internal/pgmodel"
	"net/http"
)

func TokenGet(c *gin.Context) {
	appid := c.Query("appid")
	appsecret := c.Query("appsecret")
	if appid == "" || appsecret == "" {
		c.JSON(http.StatusOK, Response.Fail("参数错误"))
		return
	}
	var thirdparty pgmodel.Thirdparty
	//pgmodel.DB.Model(&pgmodel.Thirdparty{}).Where("appid = ?", appid).Where("appsecret = ?", appsecret).Order("id ASC").Find(&thirdparty)
	pgmodel.DB.Model(&pgmodel.Thirdparty{}).Order("id DESC").Find(&thirdparty)
	c.JSON(http.StatusOK, Response.Succ(thirdparty))
	return
}
