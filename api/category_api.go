package api

import (
	"blog/global"
	"blog/models"
	"blog/res"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct {
}

// GetCategoryList 分页获取分类列表
func (CategoryApi) GetCategoryList(c *gin.Context) {
	db := global.MysqlDB
	var categoryList []models.ArticleCategory
	db.Find(&categoryList)
	res.Success(c, categoryList, "")
}
