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
	// 列表数据
	var tagsList []pgmodel.Tags
	pgmodel.DB.Model(&pgmodel.Tags{}).Where("type = $1", 0).Where("is_publish = $2", "yes").Select([]string{"name", "logo", "id"}).Order("sort ASC").Find(&tagsList)

	// 转[]map[string][string] 增加一个全路径图片地址
	var m []map[string]interface{}
	imgHost := config.YamlFile.Common.Schema + config.YamlFile.Common.Imghost
	for _, tags := range tagsList {
		item := structs.New(tags).Map()
		if val, ok := item["logo"]; ok {
			item["logo_url"] = imgHost + fmt.Sprintf("%v", val)
		}
		m = append(m, item)
	}
	c.JSON(http.StatusOK, gin.H{
		"e": 0,
		"m": "操作成功",
		"d": m,
	})
}
func HomeGetTags(c *gin.Context) {
	// 列表数据
	var tagsList []pgmodel.Tags
	pgmodel.DB.Model(&pgmodel.Tags{}).Where("type = $1", 3).Where("is_publish = $2", "yes").Select([]string{"name", "logo", "id"}).Order("sort ASC").Find(&tagsList)

	if len(tagsList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"e": 0,
			"m": "操作成功",
			"d": make([]int, 0),
		})
		return
	}
	// 转[]map[string][string] 增加一个全路径图片地址
	var m []map[string]interface{}
	imgHost := config.YamlFile.Common.Schema + config.YamlFile.Common.Imghost
	for _, tags := range tagsList {
		item := structs.New(tags).Map()
		if val, ok := item["logo"]; ok {
			item["logo_url"] = imgHost + fmt.Sprintf("%v", val)
		}
		m = append(m, item)
	}
	c.JSON(http.StatusOK, gin.H{
		"e": 0,
		"m": "操作成功",
		"d": m,
	})
}
