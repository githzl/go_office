package openController

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_office/internal/controllers"
	"go_office/internal/pgmodel"
	"go_office/internal/redis"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type responseAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

const expire = 7200

func TokenGet(c *gin.Context) {
	appid := c.Query("appid")
	appsecret := c.Query("appsecret")
	if appid == "" || appsecret == "" {
		c.JSON(http.StatusOK, controllers.Fail("缺少参数appid和appsecret"))
		return
	}
	var thirdparty pgmodel.Thirdparty
	pgmodel.DB.Model(&pgmodel.Thirdparty{}).Where("appid = ?", appid).First(&thirdparty)
	if thirdparty.Id == 0 {
		c.JSON(http.StatusOK, controllers.Fail("appid不存在"))
		return
	}
	if thirdparty.Appsecret != appsecret {
		c.JSON(http.StatusOK, controllers.Fail("appsecret错误"))
		return
	}
	if thirdparty.Enable != "yes" {
		c.JSON(http.StatusOK, controllers.Fail("该appid已经被禁用"))
		return
	}

	generateTokenErrorCount := 0
	// generateToken:
generateToken:
	str := strconv.FormatInt(time.Now().UnixNano()+rand.Int63n(10), 10)
	accessToken := fmt.Sprintf("%x", md5.Sum([]byte(str))) + fmt.Sprintf("%x", sha1.Sum([]byte(thirdparty.Appid)))

	exists, err := redis.Client.Exists(accessToken).Result()
	if err != nil {
		c.JSON(http.StatusOK, controllers.Fail(err.Error()))
		return
	}
	if exists == 1 && generateTokenErrorCount < 3 {
		generateTokenErrorCount++
		goto generateToken // 如果出现哈希碰撞 则重新生成一个access_token
	}
	if exists == 1 && generateTokenErrorCount >= 3 {
		c.JSON(http.StatusOK, controllers.Fail("生成access_token失败，请联系管理员"))
		return
	}
	if _, err := redis.Client.Set(accessToken, thirdparty.Id, time.Second*expire).Result(); err != nil {
		c.JSON(http.StatusOK, controllers.Fail(err.Error()))
		return
	}
	res := responseAccessToken{
		AccessToken: accessToken,
		ExpiresIn:   expire,
	}
	c.JSON(http.StatusOK, controllers.Succ(res))
	return
}
