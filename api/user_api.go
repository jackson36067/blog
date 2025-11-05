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
	"blog/utils"
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// GetUserDataView 获取用户信息
func (UserApi) GetUserDataView(c *gin.Context) {
	username := c.Query("username")
	db := global.MysqlDB
	var user models.User
	// 获取用户
	db.Preload("Articles").Where("username = ?", username).Take(&user)
	if user.Username == "" {
		res.Fail(c, 500, consts.UserNotFound)
	}
	// 封装用户信息
	// 获取用户粉丝以及关注数量
	var followers []uint
	var followed []uint
	db.Model(&models.UserFollow{}).Where("follower_id = ?", user.ID).Pluck("followed_id", &followed)
	db.Model(&models.UserFollow{}).Where("followed_id = ?", user.ID).Pluck("follower_id", &followers)
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
	username := c.Query("username")
	db := global.MysqlDB
	// 根据用户名称查询用户
	var userId uint
	db.Model(&models.User{}).Where("username = ?", username).Pluck("id", &userId)
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

// GetUserLikeArticlesView 获取用户的点赞博文列表
func (UserApi) GetUserLikeArticlesView(c *gin.Context) {
	db := global.MysqlDB
	// 根据用户名查询用户
	var userLikeRequestParams request.UserRequest
	var userLikeArticleIds []int
	err := c.ShouldBindQuery(&userLikeRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	// 根据用户名获取用户
	var userId uint
	db.Model(&models.User{}).Where("username = ?", userLikeRequestParams.Username).Pluck("id", &userId)
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

// GetUserBrowseArticleHistoryView 获取用户历史浏览列表
func (UserApi) GetUserBrowseArticleHistoryView(c *gin.Context) {
	username := c.Query("username")
	db := global.MysqlDB
	var userId uint
	db.Model(&models.User{}).Where("username = ?", username).Pluck("id", &userId)
	var userBrowseArticles []models.UserArticleBrowseHistory
	db.Model(&models.UserArticleBrowseHistory{}).
		Preload("Article").
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&userBrowseArticles)
	browseArticleGroup := service.GetArticleGroupedByTime(userBrowseArticles)
	res.Success(c, browseArticleGroup, "")
}

// GetUserFollowed 获取用户的关注列表
func (UserApi) GetUserFollowed(c *gin.Context) {
	var userRequestParams request.UserRequest
	err := c.ShouldBindQuery(&userRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 根据用户名查询用户id
	var userId uint
	db.Model(&models.User{}).Where("username = ?", userRequestParams.Username).Pluck("id", &userId)
	// 查询用户的关注列表
	page := userRequestParams.Page
	pageSize := userRequestParams.PageSize
	tx := db.Model(&models.UserFollow{}).Where("follower_id = ?", userId)
	// 获取关注总数量
	var total int64
	tx = tx.Count(&total)
	var userFollowed []models.UserFollow
	// 分页获取分页总数量
	tx = tx.Order("created_at desc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&userFollowed)
	userFollowedResponse := utils.MapSlice(userFollowed, func(follow models.UserFollow) response.UserFollowResponse {
		var followed models.User
		db.Where("id = ?", follow.FollowedID).Take(&followed)
		return response.UserFollowResponse{
			FollowedID: followed.ID,
			Avatar:     followed.Avatar,
			Abstract:   followed.Abstract,
			Username:   followed.Username,
			IsFollow:   true,
		}
	})
	var totalPage = int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := res.Pagination{
		Page:          page,
		PageSize:      pageSize,
		TotalElements: total,
		TotalPages:    totalPage,
		Data:          userFollowedResponse,
	}
	res.Success(c, pagination, "")
}

func (UserApi) UpdateFollow(c *gin.Context) {
	// 获取请求参数
	operateUserIdStr := c.Param("id")
	operateUserId, err := utils.StringToUint(operateUserIdStr)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	userIdAny, _ := c.Get(consts.UserId)
	userId := userIdAny.(uint)
	var userFollowRequestParam request.UserFollowRequest
	err = c.ShouldBindJSON(&userFollowRequestParam)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	isFollow := userFollowRequestParam.IsFollow
	db := global.MysqlDB
	// 返回值带上与传递是否关注的相反值, 便于前端显示
	if isFollow {
		// 取消关注 -> 从用户关注表中移除数据
		db.Where("follower_id = ? AND followed_id = ?", userId, operateUserId).Delete(&models.UserFollow{})
		res.Success(c, nil, "取关成功")
	} else {
		// 关注 -> 新增数据到用户关注表中
		db.Create(&models.UserFollow{
			FollowerID: userId,
			FollowedID: operateUserId,
		})
		res.Success(c, nil, "关注成功")
	}
}
