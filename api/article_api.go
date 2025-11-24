package api

import (
	"blog/consts"
	"blog/dto/request"
	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"
	"blog/utils"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleApi struct {
}

// GetHomeArticleView 根据条件获取文章列表
func (ArticleApi) GetHomeArticleView(c *gin.Context) {
	// 判断是游客状态还是登录状态
	authHeader := c.GetHeader("Authorization")
	var userId uint
	var isLoggedIn bool

	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if claims, err := utils.ParseToken(tokenStr); err == nil {
			userId = claims.UserID
			isLoggedIn = true
		}
	}

	// 解析请求参数
	var articleQueryParams request.ArticleQueryParams
	if err := c.ShouldBindQuery(&articleQueryParams); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	// 封装查询条件
	db := global.MysqlDB
	tx := db.Model(&models.Article{}).Preload("Category").Preload("User")
	if articleQueryParams.UserId != 0 {
		tx = tx.Where("user_id = ?", articleQueryParams.UserId)
	}
	if articleQueryParams.CategoryTitle != "" {
		var categoryId uint
		db.Model(&models.ArticleCategory{}).Where("title = ?", articleQueryParams.CategoryTitle).Pluck("id", &categoryId)
		tx = tx.Where("category_id = ?", categoryId)
	}
	if articleQueryParams.Title != "" {
		tx = tx.Where("title like ?", "%"+articleQueryParams.Title+"%")
	}
	if len(articleQueryParams.Tags) > 0 {
		for _, tag := range articleQueryParams.Tags {
			// 每个标签都要匹配，意味着文章必须包含这些标签
			tx = tx.Where("JSON_CONTAINS(tag_list, ?)", fmt.Sprintf(`"%s"`, tag))
		}
	}

	tx = tx.Where("status = ?", enum.Published)

	// 游客状态 -> 仅展示公开文章
	if !isLoggedIn {
		tx = tx.Where(
			db.
				Where("user_id = ?", userId).
				Or("visibility = ?", enum.Public),
		)
	} else { // 登录状态 -> 自己的全部 + 公开文章 + 已关注作者的粉丝文章
		// 获取我关注的作者列表
		var followedIDs []uint
		db.Model(&models.UserFollow{}).
			Where("follower_id = ?", userId).
			Pluck("followed_id", &followedIDs)
		tx = tx.
			Where(db.
				Where("user_id = ?", userId).
				Or("visibility = ?", enum.Public).
				Or("visibility = ? and user_id in ?", enum.Fans, followedIDs),
			)
	}

	page := articleQueryParams.Page
	pageSize := articleQueryParams.PageSize
	offset := (page - 1) * pageSize
	var total int64
	// 计算总元素数量
	tx.Count(&total)
	// 分页查询
	var articles []models.Article
	tx.Debug().
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&articles)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	homeArticleResponse := service.ArticlesToArticleResponse(articles)
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
	userTopArticleResponse := utils.MapSlice(userTopArticleList, func(userTopArticle models.UserTopArticle) response.ArticleResponse {
		article := userTopArticle.Article
		if article == nil {
			return response.ArticleResponse{} // 防止空指针
		}
		return response.ArticleResponse{
			Id:    article.ID,
			Title: article.Title,
		}
	})
	res.Success(c, userTopArticleResponse, "")
}

// GetArticleHotTagsAndRandCategoryView 获取10条热门文章标签以及5条随机文章分类
func (ArticleApi) GetArticleHotTagsAndRandCategoryView(c *gin.Context) {
	db := global.MysqlDB
	var articleCategoryList []models.ArticleCategory
	// 随机获取5个文章分类
	db.
		Order("RAND()").
		Offset(0).
		Limit(5).
		Find(&articleCategoryList)
	// 获取10个热门文章标签
	var hotArticleTagList []models.ArticleTag
	db.
		Order("browse_count desc").
		Offset(0).
		Limit(10).
		Find(&hotArticleTagList)
	var articleTagsAndCategoryList = response.ArticleHotTagsAndRandCategoryResponse{ArticleTags: hotArticleTagList, ArticleCategories: articleCategoryList}
	res.Success(c, articleTagsAndCategoryList, "")
}

// GetUserArticlePaginationView 分页获取用户的文章
func (ArticleApi) GetUserArticlePaginationView(c *gin.Context) {
	var myArticleQueryParam request.MyArticleQueryParams
	var userId uint
	if err := c.ShouldBindQuery(&myArticleQueryParam); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	db := global.MysqlDB
	// 根据用户名获取用户id
	db.Model(&models.User{}).Where("username = ?", myArticleQueryParam.Username).Pluck("id", &userId)
	tx := db.Model(&models.Article{}).Preload("User")
	tx = tx.Where("user_id = ?", userId)
	if myArticleQueryParam.Visibility == enum.Private {
		tx = tx.Where("visibility = ?", enum.Private)
	}
	startTime := myArticleQueryParam.StartTime
	endTime := myArticleQueryParam.EndTime
	if !startTime.IsZero() && !endTime.IsZero() {
		tx = tx.Where("created_at BETWEEN ? AND ?", startTime, endTime)
	}
	page := myArticleQueryParam.Page
	pageSize := myArticleQueryParam.PageSize
	offset := (page - 1) * pageSize
	var articleList []models.Article
	var total int64
	tx.Count(&total)
	tx.Debug().Order(fmt.Sprintf("%s %s", myArticleQueryParam.OrderBy, myArticleQueryParam.OrderType)).Offset(offset).Limit(pageSize).Find(&articleList)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	myArticleList := service.ArticlesToArticleResponse(articleList)
	pagination := res.NewPagination(page, pageSize, total, totalPages, myArticleList)
	res.Success(c, pagination, "")
}

// GetUserArticleCreateProcessView 获取文章创作历程
func (ArticleApi) GetUserArticleCreateProcessView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var articles []response.ArticleStatistic
	db.Debug().Model(&models.Article{}).
		Where("user_id = ?", userId).
		Select("id,created_at").
		Find(&articles)
	userArticleCreateProcess := service.GroupArticlesByYearAndMonth(articles)
	res.Success(c, userArticleCreateProcess, "")
}

// ClearUserBrowseArticleHistoryView 清除用户历史浏览文章记录
func (ArticleApi) ClearUserBrowseArticleHistoryView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	db.Delete(&models.UserArticleBrowseHistory{}, "user_id = ?", userId)
	res.Success(c, nil, consts.ClearUserBrowseHistorySuccess)
}

// GetArticleCategoryListView 分页获取文章分类列表
func (ArticleApi) GetArticleCategoryListView(c *gin.Context) {
	db := global.MysqlDB
	var categoryList []models.ArticleCategory
	db.Find(&categoryList)
	res.Success(c, categoryList, "")
}

// GetArticleTagListView 获取文章标签列表
func (ArticleApi) GetArticleTagListView(c *gin.Context) {
	db := global.MysqlDB
	var articleTagList []models.ArticleTag
	db.Find(&articleTagList)
	res.Success(c, articleTagList, "")
}

// GetArticleDetailView 获取文章详情信息
func (ArticleApi) GetArticleDetailView(c *gin.Context) {
	articleId := c.Param("id")
	token := c.Request.Header.Get("Authorization")
	db := global.MysqlDB
	var article models.Article
	db.Model(&models.Article{}).
		Preload("User").
		Where("id = ?", articleId).
		First(&article)
	if article.ID == 0 {
		res.Fail(c, 500, consts.ArticleNotFound)
		return
	}
	var articleResponse response.ArticleResponse
	// 访问的是无登录状态下
	if token == "" {
		articleResponse = response.ArticleResponse{
			Id:            article.ID,
			Title:         article.Title,
			Abstract:      article.Abstract,
			Content:       article.Content,
			Tags:          article.TagList,
			CreatedAt:     article.CreatedAt.Format("2006-01-02 15:04:05"),
			BrowseCount:   article.BrowseCount,
			LikeCount:     article.LikeCount,
			CollectCount:  article.CollectCount,
			CommentCount:  article.CommentCount,
			PublicComment: article.PublicComment,
			UserID:        article.UserID,
			Username:      article.User.Username,
			Avatar:        article.User.Avatar,
			IsLike:        false,
			IsCollect:     false,
			IsFollow:      false,
		}
	} else {
		token = strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
		claims, _ := utils.ParseToken(token)
		userId := claims.UserID
		// 判断登录用户是否点赞,收藏该文章
		var isLike bool
		db.Model(&models.ArticleLike{}).
			Select("1").
			Where("user_id = ? AND article_id = ?", userId, articleId).
			Limit(1).
			Find(&isLike)
		var isCollect bool
		db.Model(&models.UserArticleCollect{}).
			Select("1").
			Where("user_id = ? AND article_id = ?", userId, articleId).
			Limit(1).
			Find(&isCollect)
		var isFollow bool
		db.Model(&models.UserFollow{}).
			Select("1").
			Where("follower_id = ? AND followed_id = ?", userId, article.UserID).
			Limit(1).
			Find(&isFollow)
		articleResponse = response.ArticleResponse{
			Id:            article.ID,
			Title:         article.Title,
			Abstract:      article.Abstract,
			Content:       article.Content,
			Tags:          article.TagList,
			CreatedAt:     article.CreatedAt.Format("2006-01-02 15:04:05"),
			BrowseCount:   article.BrowseCount,
			LikeCount:     article.LikeCount,
			CollectCount:  article.CollectCount,
			CommentCount:  article.CommentCount,
			PublicComment: article.PublicComment,
			UserID:        article.UserID,
			Username:      article.User.Username,
			Avatar:        article.User.Avatar,
			IsLike:        isLike,
			IsCollect:     isCollect,
			IsFollow:      isFollow,
		}
		// 保存用户浏览文章历史记录
		go func(articleId uint, userId uint) {
			userArticleBrowseHistory := models.UserArticleBrowseHistory{
				UserID:    userId,
				ArticleID: articleId,
			}
			db.Create(&userArticleBrowseHistory)
		}(article.ID, userId)
	}
	res.Success(c, articleResponse, "")
	// 异步添加用户浏览文章记录以及文章访问量
	go func(articleId uint) {
		db.Model(&models.Article{}).Where("id = ?", articleId).Update("browse_count", gorm.Expr("browse_count + 1"))
	}(article.ID)
}

// LikeArticleView 点赞文章
func (ArticleApi) LikeArticleView(c *gin.Context) {
	// 获取点赞用户id以及文章id
	articleIdStr := c.Param("id")
	articleId, _ := strconv.ParseUint(articleIdStr, 10, 64)
	userIdAny, _ := c.Get(consts.UserId)
	userId, _ := userIdAny.(uint)
	// 获取是否点赞文章
	var likeArticleRequestParams request.LikeArticleRequestParams
	err := c.ShouldBindJSON(&likeArticleRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
		return
	}
	db := global.MysqlDB
	// 判断是点赞操作还是取消点赞操作
	if likeArticleRequestParams.IsLike {
		// 取消点赞
		db.Where("user_id = ? AND article_id = ?", userId, articleId).Delete(models.ArticleLike{})
		// 点赞数-1
		go func(articleId uint) {
			db.Model(&models.Article{}).
				Where("id = ?", articleId).
				Update("like_count", gorm.Expr("like_count - 1"))
		}(uint(articleId))
	} else {
		// 新增点赞记录
		db.Create(&models.ArticleLike{
			ArticleID: uint(articleId),
			UserID:    userId,
		})
		// 点赞数+1
		go func(articleId uint) {
			db.Model(&models.Article{}).
				Where("id = ?", articleId).
				Update("like_count", gorm.Expr("like_count + 1"))
		}(uint(articleId))
	}
	message := func() string {
		if likeArticleRequestParams.IsLike {
			return "取消点赞成功"
		}
		return "点赞成功"
	}()
	res.Success(c, nil, message)
}

// CollectArticleView 收藏文章
func (ArticleApi) CollectArticleView(c *gin.Context) {
	// 获取收藏文章id
	articleIdStr := c.Param("id")
	articleId, _ := strconv.ParseUint(articleIdStr, 10, 64)
	// 获取添加收藏夹id
	var collectArticleRequestParams request.CollectArticleRequestParams
	err := c.ShouldBindJSON(&collectArticleRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	// 获取当前收藏用户id
	userIdAny, _ := c.Get(consts.UserId)
	userId, _ := userIdAny.(uint)
	db := global.MysqlDB
	tx := db.Begin()
	// 如果该文章还没有被收藏那么就收藏数+1
	var collectRecord models.UserArticleCollect
	err = tx.Where("user_id = ? AND article_id = ?", userId, uint(articleId)).First(&collectRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := tx.Model(&models.Article{}).
			Where("id = ?", articleId).
			Update("collect_count", gorm.Expr("collect_count + 1")).
			Error; err != nil {
			tx.Rollback()
			res.Fail(c, 500, consts.CollectError)
			return
		}
	}
	// 新增收藏记录
	if err := tx.Create(&models.UserArticleCollect{
		ArticleID:  uint(articleId),
		FavoriteID: collectArticleRequestParams.FavoriteId,
		UserID:     userId,
	}).Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.CollectError)
		return
	}
	// 新增收藏夹记录
	if err := tx.Create(&models.FavoriteArticles{
		ArticleID:  uint(articleId),
		FavoriteID: collectArticleRequestParams.FavoriteId,
	}).Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.CollectError)
		return
	}
	tx.Commit()
	res.Success(c, nil, consts.CollectSuccess)
}
