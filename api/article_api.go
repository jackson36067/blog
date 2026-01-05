package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"

	"blog/consts"
	"blog/dto/request"
	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"
	"blog/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ArticleApi struct{}

// GetRecommendArticleView 获取主页推荐文章列表
func (ArticleApi) GetRecommendArticleView(c *gin.Context) {
	// 获取请求参数
	var q request.RecommendArticleQueryParams
	err := c.ShouldBindQuery(&q)
	if err != nil {
		res.Fail(c, 400, consts.RequestParamParseError)
	}
	mysqlClient := global.MysqlDB
	redisClient := global.RedisDB
	var articles []models.Article
	var totalElments int64
	// 判断用户是否登录
	if q.UserID == 0 {
		// 未登录, 返回热度高的文章(高评论, 高点赞, 高收藏综合)
		start := int64((q.Page - 1) * q.PageSize)
		stop := start + int64(q.PageSize) - 1
		articleIds, _ := redisClient.ZRevRange(global.Ctx, consts.HighQualityArticleKey, start, stop).Result()
		mysqlClient.Model(&models.Article{}).
			Preload("User").
			Where("id IN ?", articleIds).
			Order("browse_count DESC").
			Find(&articles)
		totalElments, _ = redisClient.ZCard(global.Ctx, consts.HighQualityArticleKey).Result()
	} else {
		// 登录了, 获取推荐文章列表
		articles, totalElments, err = GetRecommendedWithCacheAndPagination(mysqlClient, redisClient, q.UserID, q.HobbyTags, q.Page, q.PageSize)
		if err != nil {
			res.Fail(c, 500, consts.RequestParamParseError)
		}
	}
	articleResponse := service.ArticlesToArticleResponse(articles)
	pagination := res.Pagination{
		Page:          q.Page,
		PageSize:      q.PageSize,
		TotalElements: totalElments,
		TotalPages:    int(math.Ceil(float64(totalElments) / float64(q.PageSize))),
		Data:          articleResponse,
	}
	res.Success(c, pagination, "")
}

func GetRecommendedWithCacheAndPagination(db *gorm.DB, rdb *redis.Client, userID uint, userTags []string, page int, pageSize int) ([]models.Article, int64, error) {
	// 用户缓存用户首页文章列表
	cacheKey := fmt.Sprintf(consts.UserRecommendArticleKeyPrefix, userID)

	exists, _ := rdb.Exists(global.Ctx, cacheKey).Result()
	if exists == 0 {
		newIDs := GenerateRecommendArticleIds(db, rdb, userTags)
		rdb.Del(global.Ctx, cacheKey)
		if len(newIDs) > 0 {
			// 使用管道(Pipeline)批量写入，效率更高
			pipe := rdb.Pipeline()
			for _, id := range newIDs {
				pipe.RPush(global.Ctx, cacheKey, id)
			}
			pipe.Expire(global.Ctx, cacheKey, 1*time.Hour)
			_, err := pipe.Exec(global.Ctx)
			if err != nil {
				return nil, 0, err
			}
		}
	}

	// 2. 获取 Redis 中当前推荐池的总长度
	total, err := rdb.LLen(global.Ctx, cacheKey).Result()
	if err != nil {
		return nil, 0, err
	}

	// 3. 从 Redis 分页获取 ID 字符串列表
	start := int64((page - 1) * pageSize)
	stop := start + int64(pageSize) - 1
	idStrings, err := rdb.LRange(global.Ctx, cacheKey, start, stop).Result()

	// 如果这一页没数据，直接返回总数和空数组
	if err != nil || len(idStrings) == 0 {
		return []models.Article{}, total, nil
	}

	// 4. 根据 ID 顺序批量查询数据库详情
	var articles []models.Article
	// 保持 Redis 里的洗牌顺序
	orderBy := fmt.Sprintf("FIELD(id, %s)", strings.Join(idStrings, ","))
	err = db.Preload("User").Where("id IN ?", idStrings).Order(orderBy).Find(&articles).Error

	return articles, total, err
}

// GenerateRecommendArticleIds 获取用户感兴趣的文章以及少部分其他文章
func GenerateRecommendArticleIds(db *gorm.DB, rdb *redis.Client, userHobbyTags []string) []uint {
	var finallyIds []uint
	// 根据兴趣推荐以及其他文章比例为4:1
	recommendCount := 40
	otherCount := 10
	userHobbyTagsJSON, _ := json.Marshal(userHobbyTags)
	var recommendIds []uint
	// 获取用户感兴趣的文章, 通过用户兴趣标签匹配40篇文章
	db.Model(&models.Article{}).
		Where("JSON_OVERLAPS(tag_list, ?)", string(userHobbyTagsJSON)).
		Order("browse_count DESC"). // 根据查看人数排序
		Limit(recommendCount).
		Pluck("id", &recommendIds)
	// 在Redis中获取高质量文章十篇
	var otherIds []uint
	// 取出前四十条高质量文章进行筛选
	candidates, _ := rdb.ZRevRange(global.Ctx, consts.HighQualityArticleKey, 0, 39).Result()
	if len(candidates) > 0 {
		db.Model(&models.Article{}).
			Where("id IN ?", candidates).
			Not("JSON_OVERLAPS(tag_list, ?)", string(userHobbyTagsJSON)).
			Order("browse_count DESC"). // 根据查看人数排序
			Limit(otherCount).
			Pluck("id", &otherIds)
	}
	// 合并推荐文章和其他文章
	finallyIds = append(recommendIds, otherIds...)
	// 打乱文章顺序
	rand.Shuffle(len(finallyIds), func(i, j int) {
		finallyIds[i], finallyIds[j] = finallyIds[j], finallyIds[i]
	})
	return finallyIds
}

// GetHomeArticleView 根据条件获取文章列表
func (ArticleApi) GetHomeArticleView(c *gin.Context) {
	// 判断是游客状态还是登录状态
	currentUserId := GetUserIdFromHeader(c)
	// 解析请求参数
	var articleQueryParams request.ArticleQueryParams
	if err := c.ShouldBindQuery(&articleQueryParams); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	// 封装查询条件
	db := global.MysqlDB

	page := articleQueryParams.Page
	pageSize := articleQueryParams.PageSize
	var total int64
	// 分页查询
	var articles []models.Article

	db.Debug().
		Model(&models.Article{}).
		Preload("Category").
		Preload("User").
		Scopes(HomeArticleFilterScope(articleQueryParams), VisibilityScope(currentUserId)). // 条件过滤
		Count(&total).
		Order("created_at desc").
		Scopes(res.Paginate(page, pageSize)). // 分页
		Find(&articles)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	homeArticleResponse := service.ArticlesToArticleResponse(articles)
	res.Success(c, res.NewPagination(page, pageSize, total, totalPages, homeArticleResponse), "")
}

// GetUserIdFromHeader 从请求头中的token中解析用户id -> 不走jwt解析中间件的路由使用
func GetUserIdFromHeader(c *gin.Context) uint {
	token := c.GetHeader("Authorization")
	if token == "" {
		return 0
	}
	// 拿到后半部分
	token = strings.Split(token, "Bearer ")[1]
	// 如果token格式不对则返回false
	if token == "" {
		return 0
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		return 0
	}
	return claims.UserID
}

// HomeArticleFilterScope 处理搜索过滤
func HomeArticleFilterScope(q request.ArticleQueryParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("status = ?", enum.Published)
		if q.UserId != 0 {
			db = db.Where("user_id = ?", q.UserId)
		}
		if q.Title != "" {
			db = db.Where("title LIKE ?", "%"+q.Title+"%")
		}
		if q.CategoryTitle != "" && q.CategoryTitle != "全部" {
			// 使用子查询，减少一次代码层面的数据库交互
			db = db.Where("category_id IN (SELECT id FROM article_category WHERE title = ?)", q.CategoryTitle)
		}
		if len(q.Tags) > 0 {
			tagsJson, _ := json.Marshal(q.Tags)
			db = db.Where("JSON_OVERLAPS(tag_list, ?)", string(tagsJson)) // 只要包含这个标签的就可以
		}
		return db
	}
}

// VisibilityScope 处理复杂的可见性逻辑
func VisibilityScope(currentUserId uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if currentUserId == 0 {
			return db.Where("visibility = ?", enum.Public)
		}

		// 登录状态：自己的 OR 公开的 OR (粉丝可见且已关注)
		subQueryFollow := global.MysqlDB.Model(&models.UserFollow{}).
			Select("followed_id").
			Where("follower_id = ?", currentUserId)

		return db.Where(
			global.MysqlDB.Where("user_id = ?", currentUserId).
				Or("visibility = ?", enum.Public).
				Or("visibility = ? AND user_id IN (?)", enum.Fans, subQueryFollow),
		)
	}
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
	articleTagsAndCategoryList := response.ArticleHotTagsAndRandCategoryResponse{ArticleTags: hotArticleTagList, ArticleCategories: articleCategoryList}
	res.Success(c, articleTagsAndCategoryList, "")
}

// GetUserArticlePaginationView 分页获取用户的文章
func (ArticleApi) GetUserArticlePaginationView(c *gin.Context) {
	var q request.MyArticleQueryParams
	if err := c.ShouldBindQuery(&q); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	db := global.MysqlDB
	page := q.Page
	pageSize := q.PageSize
	var articleList []models.Article
	var total int64
	orderStr := fmt.Sprintf("%s %s", q.OrderBy, q.OrderType)
	db.Debug().
		Model(&models.Article{}).
		Preload("User").
		Scopes(UserArticleFilterScope(q)). // 条件过滤
		Count(&total).
		Order(orderStr).
		Scopes(res.Paginate(page, pageSize)). // 分页处理
		Find(&articleList)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	myArticleList := service.ArticlesToArticleResponse(articleList)
	res.Success(c, res.NewPagination(page, pageSize, total, totalPages, myArticleList), "")
}

func UserArticleFilterScope(q request.MyArticleQueryParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("user_id = (SELECT id FROM user WHERE username = ?)", q.Username)
		if q.Visibility == enum.Private {
			db = db.Where("visibility = ?", enum.Private)
		}
		if q.Status == enum.Draft {
			db = db.Where("status = ?", enum.Draft)
		} else {
			db = db.Where("status = ?", enum.Published)
		}
		startTime := q.StartTime
		endTime := q.EndTime
		if !startTime.IsZero() && !endTime.IsZero() {
			db = db.Where("created_at BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
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

// GetArticleCategoryListView 获取文章分类列表
func (ArticleApi) GetArticleCategoryListView(c *gin.Context) {
	db := global.MysqlDB
	var categoryList []models.ArticleCategory
	db.Find(&categoryList)
	res.Success(c, categoryList, "")
}

// GetArticleTagListView 获取文章标签列表
func (ArticleApi) GetArticleTagListView(c *gin.Context) {
	redis := global.RedisDB
	db := global.MysqlDB
	var articleTagList []models.ArticleTag

	val, err := redis.Get(global.Ctx, consts.ArticleTagListRedisKey).Result()

	if err == nil {
		if err := json.Unmarshal([]byte(val), &articleTagList); err == nil {
			res.Success(c, articleTagList, "")
			return
		}
	}

	if err := db.Find(&articleTagList).Error; err != nil {
		res.Fail(c, 400, "")
		return
	}

	jsonTagList, _ := json.Marshal(articleTagList)
	expiration := 3 * 24 * time.Hour

	// 使用 Set 操作存储完整 JSON 串
	err = redis.Set(global.Ctx, consts.ArticleTagListRedisKey, jsonTagList, expiration).Err()
	if err != nil {
		log.Fatalf("redis写入失败")
	}
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
		Preload("Category").
		Where("id = ?", articleId).
		First(&article)
	if article.ID == 0 {
		res.Fail(c, 500, consts.ArticleNotFound)
		return
	}
	// 获取文章象征性评论(1条热评)
	var articleSymbolComment models.Comment
	db.Preload("User").
		Where("article_id = ? AND root_parent_id IS NULL", articleId).
		Order("like_count DESC").
		Limit(1).
		Find(&articleSymbolComment)
	var commentResponse *response.CommentResponse
	if articleSymbolComment.ID > 0 {
		commentResponse = &response.CommentResponse{
			ID:       articleSymbolComment.ID,
			UserID:   articleSymbolComment.UserID,
			Username: articleSymbolComment.User.Username,
			Avatar:   articleSymbolComment.User.Avatar,
			Content:  articleSymbolComment.Content,
		}
	}
	var articleTotalComments int64
	var articleTotalRootComment int64
	// 统计所有评论数量（包含子评论）
	db.Model(&models.Comment{}).
		Where("article_id = ?", articleId).
		Count(&articleTotalComments)
	// 统计根评论数量（root_parent_id = NULL）
	db.Model(&models.Comment{}).
		Where("article_id = ? AND root_parent_id IS NULL", articleId).
		Count(&articleTotalRootComment)
	var articleResponse response.ArticleResponse
	// 访问的是无登录状态下
	if token == "" {
		articleResponse = service.UnifyArticleDetailResult(article,
			false,
			false,
			false,
			commentResponse,
			articleTotalComments,
			articleTotalRootComment,
		)
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
		articleResponse = service.UnifyArticleDetailResult(
			article,
			isLike,
			isCollect,
			isFollow,
			commentResponse,
			articleTotalComments,
			articleTotalRootComment,
		)
		// 保存用户浏览文章历史记录
		go func(articleId uint, userId uint) {
			// 查看用户今日是否访问过了
			var userArticleBrowseHistory models.UserArticleBrowseHistory
			// 今天开始时间
			startOfDay := time.Date(
				time.Now().Year(),
				time.Now().Month(),
				time.Now().Day(),
				0, 0, 0, 0,
				time.Local,
			)
			// 今天结束时间
			endOfDay := startOfDay.Add(24 * time.Hour)
			// 查询是否存在记录
			db.Where(
				"article_id = ? AND user_id = ? AND created_at BETWEEN ? AND ?",
				articleId, userId, startOfDay, endOfDay,
			).Find(&userArticleBrowseHistory)
			if userArticleBrowseHistory.UserID > 0 {
				// 存在,修改浏览时间
				db.Model(&userArticleBrowseHistory).
					Update("created_at", time.Now())
			} else {
				userArticleBrowseHistory = models.UserArticleBrowseHistory{
					UserID:    userId,
					ArticleID: articleId,
				}
				db.Create(&userArticleBrowseHistory)
			}
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

// GetArticleCommentsByPaginationView 分页获取文章评论列表
func (ArticleApi) GetArticleCommentsByPaginationView(c *gin.Context) {
	articleIdStr := c.Param("id")
	articleId, _ := utils.StringToUint(articleIdStr)
	token := c.Request.Header.Get("Authorization")
	var params request.ArticleCommentRequestParams
	err := c.ShouldBindQuery(&params)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	var rootComments []models.Comment
	db.Model(&models.Comment{}).
		Preload("User").
		Where("article_id = ? AND root_parent_id IS NULL", articleId).
		Order("created_at DESC").
		Scopes(res.Paginate(params.Page, params.PageSize)).
		Find(&rootComments)
	// 抽取 rootIDs
	rootIDs := make([]uint, 0, len(rootComments))
	for _, r := range rootComments {
		rootIDs = append(rootIDs, r.ID)
	}
	commentResponses := service.GetArticleComments(db, rootIDs, rootComments, token)
	res.Success(c, commentResponses, "")
}

// PublishArticleView 发布文章
func (ArticleApi) PublishArticleView(c *gin.Context) {
	userIdAny, _ := c.Get(consts.UserId)
	userId, _ := userIdAny.(uint)
	var publishArticleRequestParams request.PublishArticleRequestParams
	err := c.ShouldBindJSON(&publishArticleRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 判断该文章分类是否存在,不存在创建出来
	err = db.Transaction(func(tx *gorm.DB) error {
		var articleCategory models.ArticleCategory
		// 查找不存在就会新增,新增后会回显ID
		err := db.Where("LOWER(title) = LOWER(?)", publishArticleRequestParams.CategoryName).
			FirstOrCreate(&articleCategory, models.ArticleCategory{
				Title: publishArticleRequestParams.CategoryName,
			}).Error // 新增文章信息
		if err != nil {
			res.Fail(c, 500, consts.UpdateCategoryError)
			return err
		}
		article := models.Article{
			Title:         publishArticleRequestParams.Title,
			Abstract:      publishArticleRequestParams.Abstract,
			Content:       publishArticleRequestParams.Content,
			UserID:        userId,
			CategoryID:    articleCategory.ID,
			Coverage:      publishArticleRequestParams.Coverage,
			Visibility:    publishArticleRequestParams.Visibility,
			TagList:       publishArticleRequestParams.Tags,
			BrowseCount:   0,
			LikeCount:     0,
			CommentCount:  0,
			CollectCount:  0,
			PublicComment: publishArticleRequestParams.PublicComment,
			Status:        publishArticleRequestParams.Status,
		}
		err = db.Create(&article).Error
		if err != nil {
			res.Fail(c, 500, consts.PublishArticleError)
			return err
		}
		// 返回文章id用于跳转页面
		res.Success(c, article.ID, consts.PublishArticleSuccess)
		return nil
	})
	if err != nil {
		res.Fail(c, 500, err.Error())
	}
}

// UpdateArticleView 更新文章信息
func (ArticleApi) UpdateArticleView(c *gin.Context) {
	// 解析参数
	articleId := c.Param("id")
	var publishArticleRequestParams request.PublishArticleRequestParams
	err := c.ShouldBindJSON(&publishArticleRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	tx := db.Begin()
	var articleCategory models.ArticleCategory
	err = tx.Where("LOWER(title) = LOWER(?)", publishArticleRequestParams.CategoryName).FirstOrCreate(&articleCategory, models.ArticleCategory{
		Title: publishArticleRequestParams.CategoryName,
	}).Error
	if err != nil {
		res.Fail(c, 400, consts.UpdateCategoryError)
		tx.Rollback()
	}
	tagJSON, err := json.Marshal(publishArticleRequestParams.Tags)
	if err != nil {
		res.Fail(c, 400, consts.UpdateArticleError)
		tx.Rollback()
	}
	err = tx.Model(&models.Article{}).
		Where("id = ?", articleId).
		Updates(map[string]any{
			"title":          publishArticleRequestParams.Title,
			"content":        publishArticleRequestParams.Content,
			"abstract":       publishArticleRequestParams.Abstract,
			"coverage":       publishArticleRequestParams.Coverage,
			"tag_list":       string(tagJSON),
			"public_comment": publishArticleRequestParams.PublicComment,
			"category_id":    articleCategory.ID,
			"visibility":     publishArticleRequestParams.Visibility,
			"status":         publishArticleRequestParams.Status,
		}).
		Error
	if err != nil {
		res.Fail(c, 400, consts.UpdateArticleError)
		tx.Rollback()
	}
	tx.Commit()
}
