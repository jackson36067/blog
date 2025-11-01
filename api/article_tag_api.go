package api

import (
	"blog/global"
	"blog/models"
	"blog/res"

	"github.com/gin-gonic/gin"
)

type ArticleTagApi struct{}

// GetHotArticleTagPagination 获取前十热门的文章标签
func (ArticleTagApi) GetHotArticleTagPagination(c *gin.Context) {
	db := global.MysqlDB
	var hotArticleTagList []models.ArticleTag
	db.
		Order("browse_count desc").
		Offset(0).
		Limit(10).
		Find(&hotArticleTagList)
	res.Success(c, hotArticleTagList, "")
}
