package api

import (
	"blog/consts"
	"blog/core"
	"blog/dto/request"
	"blog/dto/response"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// GetUserDataView 获取用户信息
func (UserApi) GetUserDataView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var user models.User
	// 获取用户
	db.Preload("Articles").Where("id = ?", userId).Take(&user)
	if user.Username == "" {
		res.Fail(c, 500, consts.UserNotFound)
	}
	// 封装用户信息
	// 获取用户粉丝以及关注数量
	var followers []uint
	var followed []uint
	db.Model(&models.UserFollow{}).Where("follower_id = ?", userId).Pluck("followed_id", &followed)
	db.Model(&models.UserFollow{}).Where("followed_id = ?", userId).Pluck("follower_id", &followers)
	// 获取ip地址
	core.InitIPDB()
	ip, err := core.GetIpAddress(c.ClientIP())
	if err != nil {
		res.Fail(c, 500, consts.IpParseError)
	}
	userDataResponse := response.UserDataResponse{
		OriginArticle: len(user.Articles),
		Fans:          len(followers),
		Follow:        len(followed),
		IP:            ip,
		JoinTime:      user.CreatedAt.Format("2006-01-02 15:04:05"),
		CodeAge:       user.CodeAge,
		Avatar:        user.Avatar,
	}
	res.Success(c, userDataResponse, "")
}

// GetUserAchievementListView 获取用户的成就
func (UserApi) GetUserAchievementListView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	// 获取用户所有文章获取的总点赞次数
	var userArticleIds []int
	// 获取用户所有文章
	db.Model(&models.Article{}).Where("user_id = ?", userId).Pluck("id", &userArticleIds)
	var totalLikes int64
	var totalCollects int64
	var totalComments int64
	// 获取用户所有文章的总共点赞
	db.Model(&models.ArticleLike{}).Where("article_id in (?)", userArticleIds).Count(&totalLikes)
	// 获取用户所有文章被收藏总数
	db.Model(&models.UserArticleCollect{}).Where("article_id in (?)", userArticleIds).Count(&totalCollects)
	// 获取用户所有文章的总评论数量
	db.Model(&models.Comment{}).Where("article_id in (?)", userArticleIds).Count(&totalComments)
	userAchievement := response.UserAchievementResponse{
		TotalLikes:    totalLikes,
		TotalCollects: totalCollects,
		TotalComments: totalComments,
	}
	res.Success(c, userAchievement, "")
}

func (UserApi) GetUserLikeArticlesView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var userLikeRequestParams request.UserLikesRequest
	var userLikeArticleIds []int
	err := c.ShouldBindQuery(&userLikeRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db.Model(&models.ArticleLike{}).
		Where("user_id = ?", userId).
		Order("created_at desc").
		Pluck("article_id", &userLikeArticleIds)
	var articles []models.Article
	page := userLikeRequestParams.Page
	pageSize := userLikeRequestParams.PageSize
	tx := db.Model(&models.Article{})
	tx = tx.Where("id in (?)", userLikeArticleIds)
	var total int64
	tx = tx.Count(&total)
	tx = tx.
		Order(fmt.Sprintf("FIELD(id, %s)", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(userLikeArticleIds)), ","), "[]"))).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&articles)
	articleResponse := service.ArticlesToArticleResponse(articles)
	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := res.Pagination{
		Page:          page,
		PageSize:      pageSize,
		TotalElements: total,
		TotalPages:    totalPage,
		Data:          articleResponse,
	}
	res.Success(c, pagination, "")
}

func (UserApi) GetUserBrowseArticleHistoryView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var userBrowseArticles []models.UserArticleBrowseHistory
	db.Model(&models.UserArticleBrowseHistory{}).
		Preload("Article").
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&userBrowseArticles)
	browseArticleGroup := service.GetArticleGroupedByTime(userBrowseArticles)
	res.Success(c, browseArticleGroup, "")
}
