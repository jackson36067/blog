package api

import (
	"blog/consts"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/utils"
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

type ArticleQueryParams struct {
	Page       int      `form:"page" binding:"required"`
	PageSize   int      `form:"pageSize" binding:"required"`
	Title      string   `form:"title"`      // 文章标题
	CategoryId uint     `form:"categoryId"` // 文章分类
	Tags       []string `form:"tags"`       // 文章标签
	UserId     uint     `form:"userId"`     // 文章所属用户
}

type ArticleResponse struct {
	Id            uint     `json:"id"` // 文章id
	Title         string   `json:"title"`
	Abstract      string   `json:"abstract"`
	Content       string   `json:"content"`
	Coverage      string   `json:"coverage"`
	Tags          []string `json:"tags"`
	CreatedAt     string   `json:"createdAt"`
	BrowseCount   int      `json:"browseCount"`
	LikeCount     int      `json:"likeCount"`
	CommentCount  int      `json:"commentCount"`
	CollectCount  int      `json:"collectCount"`
	PublicComment bool     `json:"publicComment"`
}

// GetHomeArticleView 根据条件获取文章列表
func (ArticleApi) GetHomeArticleView(c *gin.Context) {
	// 解析请求参数
	var homeArticleQueryParams ArticleQueryParams
	if err := c.ShouldBindQuery(&homeArticleQueryParams); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	// 封装查询条件
	db := global.MysqlDB
	tx := db.Model(&models.Article{}).Preload("Category").Preload("User")
	if homeArticleQueryParams.UserId != 0 {
		tx = tx.Where("user_id = ?", homeArticleQueryParams.UserId)
	}
	if homeArticleQueryParams.CategoryId != 0 {
		tx = tx.Where("category_id = ?", homeArticleQueryParams.CategoryId)
	}
	if homeArticleQueryParams.Title != "" {
		tx = tx.Where("title like ?", "%"+homeArticleQueryParams.Title+"%")
	}
	if len(homeArticleQueryParams.Tags) > 0 {
		for _, tag := range homeArticleQueryParams.Tags {
			// 每个标签都要匹配，意味着文章必须包含这些标签
			tx = tx.Where("JSON_CONTAINS(tag_list, ?)", fmt.Sprintf(`"%s"`, tag))
		}
	}
	tx = tx.Where("status = ?", enum.Published)
	page := homeArticleQueryParams.Page
	pageSize := homeArticleQueryParams.PageSize
	offset := (page - 1) * pageSize
	var total int64
	// 计算总元素数量
	tx.Count(&total)
	// 分页查询
	var articles []models.Article
	tx.Order("created_at desc").Offset(offset).Limit(pageSize).Find(&articles)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	homeArticleResponse := utils.MapSlice(articles, func(a models.Article) ArticleResponse {
		return ArticleResponse{
			Id:            a.ID,
			Title:         a.Title,
			Abstract:      a.Abstract,
			Content:       a.Content,
			Coverage:      a.Coverage,
			Tags:          a.TagList,
			BrowseCount:   a.BrowseCount,
			LikeCount:     a.LikeCount,
			CommentCount:  a.CommentCount,
			CollectCount:  a.CollectCount,
			CreatedAt:     a.CreatedAt.Format("2006-01-02 15:04:05"),
			PublicComment: a.PublicComment,
		}
	})
	pagination := res.NewPagination(page, pageSize, total, totalPages, homeArticleResponse)
	res.Success(c, pagination, "")
}

// GetUserTopArticleListView 获取用户置顶文章
func (ArticleApi) GetUserTopArticleListView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var userTopArticleList []models.UserTopArticle
	db.Preload("Article").
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&userTopArticleList)
	userTopArticleResponse := utils.MapSlice(userTopArticleList, func(userTopArticle models.UserTopArticle) ArticleResponse {
		article := userTopArticle.Article
		if article == nil {
			return ArticleResponse{} // 防止空指针
		}
		return ArticleResponse{
			Id:    article.ID,
			Title: article.Title,
		}
	})
	res.Success(c, userTopArticleResponse, "")
}
