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
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserApi struct{}

// GetUserDataView 获取用户信息
func (UserApi) GetUserDataView(c *gin.Context) {
	// 传递用户名为了获取他人的信息
	username := c.Query("username")
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var user models.User
	// 获取用户
	if username != "" {
		db.Model(&models.User{}).Preload("Articles").Where("username = ?", username).Find(&user)
	} else {
		db.Model(&models.User{}).Preload("Articles").Where("id = ?", userId).Find(&user)
	}
	if user.Username == "" {
		res.Fail(c, 500, consts.UserNotFound)
		return
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
		return
	}
	// 获取用户配置信息
	var userConfig models.UserConfig
	db.Where("user_id=?", user.ID).Find(&userConfig)
	userDataResponse := response.UserDataResponse{
		OriginArticle:               len(user.Articles),
		Fans:                        len(followers),
		Follow:                      len(followed),
		IP:                          ip,
		JoinTime:                    user.CreatedAt.Format("2006-01-02 15:04:05"),
		CodeAge:                     user.CodeAge,
		Username:                    user.Username,
		Avatar:                      user.Avatar,
		Sex:                         user.Sex,
		Abstract:                    user.Abstract,
		Birthday:                    user.Birthday.Format("2006-01-02"),
		HobbyTags:                   userConfig.HobbyTags,
		PublicFanList:               userConfig.PublicFanList,
		PublicCollectList:           userConfig.PublicCollectList,
		PublicFollowList:            userConfig.PublicFollowList,
		SinceLastUpdateUsernameDays: int(time.Since(userConfig.UpdateUsernameDate).Hours() / 24),
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
	var userLikeRequestParams request.UserRequestParams
	err := c.ShouldBindQuery(&userLikeRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 根据用户名获取用户
	var userId uint
	db.Model(&models.User{}).
		Where("username = ?", userLikeRequestParams.Username).
		Pluck("id", &userId)
	// 获取用户点赞的文章id
	var userLikeArticleIds []uint
	db.Model(&models.ArticleLike{}).
		Where("user_id = ?", userId).
		Order("created_at desc").
		Pluck("article_id", &userLikeArticleIds)
	var articles []models.Article
	page := userLikeRequestParams.Page
	pageSize := userLikeRequestParams.PageSize
	var total int64
	db.Model(&models.Article{}).
		Where("id in (?)", userLikeArticleIds).
		Count(&total).
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
	var userRequestParams request.UserRequestParams
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
			ID:       followed.ID,
			Avatar:   followed.Avatar,
			Abstract: followed.Abstract,
			Username: followed.Username,
			IsFollow: true,
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
	var userFollowRequestParam request.UserFollowRequestParam
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

// GetUserFollower 分页获取用户粉丝列表
func (UserApi) GetUserFollower(c *gin.Context) {
	var userRequestParams request.UserRequestParams
	err := c.ShouldBindQuery(&userRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 根据用户名查询用户id
	var userId uint
	db.Model(&models.User{}).Where("username = ?", userRequestParams.Username).Pluck("id", &userId)
	// 查询用户的粉丝列表
	page := userRequestParams.Page
	pageSize := userRequestParams.PageSize
	tx := db.Model(&models.UserFollow{}).Where("followed_id = ?", userId)
	// 获取粉丝总数量
	var total int64
	tx = tx.Count(&total)
	var userFollower []models.UserFollow
	// 分页获取粉丝总数量
	tx = tx.Order("created_at desc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&userFollower)
	userFollowerResponse := utils.MapSlice(userFollower, func(follow models.UserFollow) response.UserFollowResponse {
		var follower models.User
		db.Where("id = ?", follow.FollowerID).Take(&follower)
		// 判断用户是否关注其粉丝
		var followed models.UserFollow
		db.Where("follower_id = ? And followed_id = ?", userId, follower.ID).Find(&followed)
		return response.UserFollowResponse{
			ID:       follower.ID,
			Avatar:   follower.Avatar,
			Abstract: follower.Abstract,
			Username: follower.Username,
			IsFollow: followed.ID > 0,
		}
	})
	var totalPage = int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := res.Pagination{
		Page:          page,
		PageSize:      pageSize,
		TotalElements: total,
		TotalPages:    totalPage,
		Data:          userFollowerResponse,
	}
	res.Success(c, pagination, "")
}

// GetUserComments 分页获取用户发布/收到的评论
func (UserApi) GetUserComments(c *gin.Context) {
	var userCommentRequestParams request.UserCommentRequestParams
	err := c.ShouldBindQuery(&userCommentRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 根据用户名查询用户id
	var userId uint
	db.Model(&models.User{}).
		Where("username = ?", userCommentRequestParams.Username).
		Pluck("id", &userId)
	page := userCommentRequestParams.Page
	pageSize := userCommentRequestParams.PageSize
	// 判断获取自己发布的评论还是收到的评论
	var total int64
	var comments []models.Comment
	switch userCommentRequestParams.Type {
	case "in":
		// 获取收到的评论
		var articleIds []uint
		// 获取文章并携带评论, 评论根据创建时间排序
		db.Model(&models.Article{}).Where("user_id = ?", userId).Pluck("id", &articleIds)
		if len(articleIds) == 0 {
			comments = make([]models.Comment, 0)
		} else {
			db.Model(&models.Comment{}).
				Where("article_id IN (?)", articleIds).
				Preload("Article").
				Order("created_at DESC").
				Count(&total).
				Offset((page - 1) * pageSize).
				Limit(pageSize).
				Find(&comments)
		}
	case "out":
		// 获取自己发布的评论
		db.Model(&models.Comment{}).
			Where("user_id = ?", userId).
			Preload("Article").
			Order("created_at DESC").
			//Count(&total).
			Offset((page - 1) * pageSize).
			Limit(pageSize).
			Find(&comments)
	default:
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))
	userCommentResponse := utils.MapSlice(comments, func(comment models.Comment) response.UserCommentResponse {
		return response.UserCommentResponse{
			CommentID: comment.ID,
			ArticleID: comment.ArticleID,
			Content:   comment.Content,
			Title:     comment.Article.Title,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	})
	pagination := res.Pagination{
		Page:          page,
		PageSize:      pageSize,
		TotalElements: total,
		TotalPages:    totalPage,
		Data:          userCommentResponse,
	}
	res.Success(c, pagination, "")
}

// Upload 文件上传
func (UserApi) Upload(c *gin.Context) {
	// 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		res.Fail(c, 500, consts.FileParseError)
		return
	}
	fileURL, err := utils.UploadImage(file)
	if err != nil {
		res.Fail(c, 500, consts.UploadFileError)
		return
	}
	res.Success(c, fileURL, consts.UploadFileSuccess)
}

// UpdateUserInfo 更改用户信息
func (UserApi) UpdateUserInfo(c *gin.Context) {
	// 解析参数
	var updateUserRequestParams request.UpdateUserRequestParams
	err := c.ShouldBindJSON(&updateUserRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
		return
	}
	db := global.MysqlDB
	// 存储更改信息
	var updateUserInfoMap = make(map[string]any)
	// 存储用户配置更改信息
	var updateUserConfigMap = make(map[string]any)
	tx := db.Begin()
	if updateUserRequestParams.Username != "" {
		// 判断是否满足更改用户名条件
		var userConfig models.UserConfig
		var user models.User
		tx.Where("user_id = ?", updateUserRequestParams.UserID).Find(&userConfig)
		tx.Where("username = ?", updateUserRequestParams.Username).Find(&user)
		if user.ID > 0 {
			res.Fail(c, 500, consts.UsernameExist)
			tx.Rollback()
			return
		}
		differ := time.Now().Sub(userConfig.UpdateUsernameDate)
		// 判断上一次改名是否超过30天
		if differ > 30*24*time.Hour {
			updateUserInfoMap["username"] = updateUserRequestParams.Username
			updateUserConfigMap["update_username_date"] = time.Now()
		}
	}
	if updateUserRequestParams.Sex != nil {
		updateUserInfoMap["sex"] = *updateUserRequestParams.Sex
	}
	if updateUserRequestParams.Avatar != "" {
		updateUserInfoMap["avatar"] = updateUserRequestParams.Avatar
	}
	if updateUserRequestParams.Birthday != "" {
		layout := "2006-01-02"
		birthday, err := time.ParseInLocation(layout, updateUserRequestParams.Birthday, time.Local)
		if err != nil {
			res.Fail(c, 500, "生日格式错误")
			return
		}
		updateUserInfoMap["birthday"] = birthday
	}
	if updateUserRequestParams.Abstract != "" {
		updateUserInfoMap["abstract"] = updateUserRequestParams.Abstract
	}
	if updateUserRequestParams.Email != "" {
		var user models.User
		db.Where("email = ?", updateUserRequestParams.Email).Find(&user)
		if user.ID > 0 {
			res.Fail(c, 500, consts.EmailExist)
			return
		}
		updateUserInfoMap["email"] = updateUserRequestParams.Email
	}
	if updateUserRequestParams.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(updateUserRequestParams.Password), bcrypt.DefaultCost)
		updateUserInfoMap["password"] = string(hash)
	}
	if updateUserRequestParams.PublicCollectList != nil {
		updateUserConfigMap["public_collect_list"] = *updateUserRequestParams.PublicCollectList
	}
	if updateUserRequestParams.PublicFollowList != nil {
		updateUserConfigMap["public_follow_list"] = *updateUserRequestParams.PublicFollowList
	}
	if updateUserRequestParams.PublicFanList != nil {
		updateUserConfigMap["public_fan_list"] = *updateUserRequestParams.PublicFanList
	}
	if updateUserRequestParams.HobbyTags != nil {
		marshal, _ := json.Marshal(updateUserRequestParams.HobbyTags)
		updateUserConfigMap["hobby_tags"] = marshal
	}
	if len(updateUserInfoMap) > 0 {
		if err := tx.Model(&models.User{}).
			Where("id = ?", updateUserRequestParams.UserID).
			Updates(updateUserInfoMap).Error; err != nil {
			res.Fail(c, 500, consts.UpdateUserError)
			tx.Rollback()
			return
		}
	}
	if len(updateUserConfigMap) > 0 {
		if err := tx.Model(&models.UserConfig{}).
			Where("user_id = ?", updateUserRequestParams.UserID).
			Updates(updateUserConfigMap).Error; err != nil {
			res.Fail(c, 500, consts.UpdateUserConfigError)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}
