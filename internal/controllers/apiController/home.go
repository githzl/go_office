package apiController

import (
	"github.com/gin-gonic/gin"
	"go_office/config"
	"go_office/internal/controllers"
	"go_office/internal/pgmodel"
	"go_office/internal/services/dto"
	"net/http"
)

func HomeGetTags(c *gin.Context) {
	// 列表数据
	var tagsList []pgmodel.Tags
	pgmodel.DB.Model(&pgmodel.Tags{}).Where("type = ?", 0).Where("is_publish = ?", "yes").Select("name,id,logo").Order("sort ASC").Find(&tagsList)
	if len(tagsList) == 0 {
		c.JSON(http.StatusOK, controllers.Succ(make([]int, 0)))
		return
	}

	imgHost := config.YamlFile.Common.Schema + config.YamlFile.Common.Imghost
	// 数据转化 返回值增加其他参数
	tagsDtoList := make([]dto.Tags, len(tagsList))
	for idx, tags := range tagsList {
		tagsDtoList[idx] = dto.Tags{
			Tags: pgmodel.Tags{
				Id:   tags.Id,
				Name: tags.Name,
				Logo: tags.Logo,
			},
			LogoUrl: imgHost + tags.Logo}
	}
	c.JSON(http.StatusOK, controllers.Succ(tagsDtoList))
	return
}
