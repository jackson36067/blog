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
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	tx.Debug().Order("created_at desc").Offset(offset).Limit(pageSize).Find(&articles)
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
	tx := db.Model(&models.Article{})
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

// GetUserArticleCreateProcess 获取文章创作历程
func (ArticleApi) GetUserArticleCreateProcess(c *gin.Context) {
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
