package apiController

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_office/config"
	"go_office/internal/pgmodel"
	_ "go_office/internal/pgmodel"
	"net/http"
)

func HomeSiteOption(c *gin.Context) {
	
}
func HomeGetTags(c *gin.Context) {
	// 列表数据
	var tagsList []pgmodel.Tags
	pgmodel.DB.Model(&pgmodel.Tags{}).Where("type = ?", 10).Where("is_publish = ?", "yes").Select([]string{"name", "logo", "id"}).Order("sort ASC").Find(&tagsList)

	res := gin.H{"e": 0, "m": "操作成功", "d": make([]int, 0)}
	if len(tagsList) == 0 {
		c.JSON(http.StatusOK, res)
		return
	}
	// 转[]map[string]string 返回值增加一个图片全路径地址字段logo_url
	var m []map[string]interface{}
	imgHost := config.YamlFile.Common.Schema + config.YamlFile.Common.Imghost
	for _, tags := range tagsList {
		item := structs.New(tags).Map()
		if val, ok := item["logo"]; ok {
			item["logo_url"] = imgHost + fmt.Sprintf("%v", val)
		}
		m = append(m, item)
	}
	if len(m) != 0 {
		res["d"] = m
	}
	c.JSON(http.StatusOK, res)
	return
}
