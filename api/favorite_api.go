package api

import (
	"blog/consts"
	"blog/dto/request"
	"blog/dto/response"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteApi struct{}

// GetUserFavoriteListView 获取用户收藏夹列表
func (FavoriteApi) GetUserFavoriteListView(c *gin.Context) {
	username := c.Query("username")
	db := global.MysqlDB
	// 根据用户名获取用户id
	var userId uint
	db.Model(&models.User{}).Where("username = ?", username).Pluck("id", &userId)
	var favoriteList []models.Favorite
	db.Preload("Articles").Where("user_id = ?", userId).Find(&favoriteList)
	favoriteListResponse := utils.MapSlice(favoriteList, func(f models.Favorite) response.FavoriteListResponse {
		return response.FavoriteListResponse{
			ID:           f.ID,
			Title:        f.Title,
			Abstract:     f.Abstract,
			IsDefault:    f.IsDefault,
			ArticleCount: len(f.Articles),
		}
	})
	res.Success(c, favoriteListResponse, "")
}

// NewFavoriteView 创建收藏夹
func (FavoriteApi) NewFavoriteView(c *gin.Context) {
	userIdAny, _ := c.Get(consts.UserId)
	// 断言用户id类型
	userId, ok := userIdAny.(uint)
	if !ok {
		res.Fail(c, 500, consts.UnKnowUserIdType)
	}
	var newFavoriteRequestParams request.NewFavoriteRequestParams
	err := c.ShouldBindJSON(&newFavoriteRequestParams)
	if err != nil {
		res.Fail(c, http.StatusBadRequest, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 开启事务
	tx := db.Begin()

	// 出现错误回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if newFavoriteRequestParams.IsDefault {
		var favorite models.Favorite
		err := tx.Where("user_id = ? AND is_default = ?", userId, true).Find(&favorite).Error
		if err != nil {
			tx.Rollback()
		}
		if favorite.ID > 0 {
			tx.Model(&models.Favorite{}).Where("id = ?", favorite.ID).Updates(map[string]interface{}{"is_default": false})
		}
	}
	// 创建新的收藏夹
	if err := tx.Create(&models.Favorite{
		UserID:    userId,
		Title:     newFavoriteRequestParams.Title,
		Abstract:  newFavoriteRequestParams.Abstract,
		IsDefault: newFavoriteRequestParams.IsDefault,
	}).Error; err != nil {
		tx.Rollback()
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		res.Fail(c, http.StatusInternalServerError, consts.AffairCommitError)
	}
	res.Success(c, nil, consts.NewFavoriteSuccess)
}

// GetFavoriteArticleListView 获取收藏夹博文列表
func (FavoriteApi) GetFavoriteArticleListView(c *gin.Context) {
	// 从路径参数中获取收藏夹id
	favoriteId := c.Param("id")
	db := global.MysqlDB
	var articles []models.Article
	err := db.Table("article").
		Joins("JOIN favorite_articles fa ON fa.article_id = article.id").
		Where("fa.favorite_id = ?", favoriteId).
		Order("fa.created_at DESC").
		Find(&articles).Error
	if err != nil {
		res.Success(c, make(map[string]any), "")
		return
	}
	favoriteArticleResponse := utils.MapSlice(articles, func(article models.Article) response.FavoriteArticleResponse {
		return response.FavoriteArticleResponse{
			ID:    article.ID,
			Title: article.Title,
		}
	})
	res.Success(c, favoriteArticleResponse, "")
}

// UpdateFavorite 更新收藏夹基本信息
func (FavoriteApi) UpdateFavorite(c *gin.Context) {
	favoriteId := c.Param("id")
	userId, _ := c.Get(consts.UserId)
	// 获取请求参数
	var updateFavoriteRequestParams request.UpdateFavoriteRequestParams
	err := c.ShouldBindJSON(&updateFavoriteRequestParams)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	tx := global.MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	updates := map[string]interface{}{}
	if updateFavoriteRequestParams.IsDefault {
		var favorite models.Favorite
		err := tx.Where("user_id = ? AND is_default = ?", userId, true).Find(&favorite).Error
		if err != nil {
			tx.Rollback()
		}
		if favorite.ID > 0 {
			tx.Model(&models.Favorite{}).Where("id = ?", favorite.ID).Updates(map[string]interface{}{"is_default": false})
		}
		updates["is_default"] = true
	}
	if updateFavoriteRequestParams.Title != "" {
		updates["title"] = updateFavoriteRequestParams.Title
	}
	if updateFavoriteRequestParams.Abstract != "" {
		updates["abstract"] = updateFavoriteRequestParams.Abstract
	}
	if len(updates) == 0 {
		res.Fail(c, 500, consts.NoUpdateField)
		tx.Rollback()
		return
	}
	tx.Model(&models.Favorite{}).Where("id = ?", favoriteId).Updates(updates)
	tx.Commit()
	res.Success(c, nil, consts.UpdateSuccess)
}

// MoveFavoriteArticleView 移动收藏夹博文至其他收藏夹
func (FavoriteApi) MoveFavoriteArticleView(c *gin.Context) {
	var moveFavoriteRequestParams request.MoveFavoriteRequestParams
	if err := c.ShouldBindJSON(&moveFavoriteRequestParams); err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
		return
	}

	db := global.MysqlDB
	tx := db.Begin()

	// 1. 查询目标收藏夹已有的文章
	var existingArticleIDs []uint
	if err := tx.Model(&models.FavoriteArticles{}).
		Where("favorite_id = ?", moveFavoriteRequestParams.TargetFavoriteID).
		Pluck("article_id", &existingArticleIDs).Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.FindTargetFavoriteError)
		return
	}

	// 2. 构造待插入的新文章列表（过滤已存在的）
	existingSet := make(map[uint]bool)
	for _, id := range existingArticleIDs {
		existingSet[id] = true
	}

	var toInsert []models.FavoriteArticles
	for _, articleId := range moveFavoriteRequestParams.ArticleIDs {
		if !existingSet[articleId] {
			toInsert = append(toInsert, models.FavoriteArticles{
				FavoriteID: moveFavoriteRequestParams.TargetFavoriteID,
				ArticleID:  articleId,
			})
		}
	}

	// 3. 批量插入新记录
	if len(toInsert) > 0 {
		if err := tx.Create(&toInsert).Error; err != nil {
			tx.Rollback()
			res.Fail(c, 500, consts.InsertNewRecordError)
			return
		}
	}

	// 4. 删除来源收藏夹中的对应文章（仅限移动的那些）
	if err := tx.
		Where("favorite_id = ? AND article_id IN ?", moveFavoriteRequestParams.SourceFavoriteID, moveFavoriteRequestParams.ArticleIDs).
		Delete(&models.FavoriteArticles{}).Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.DeleteTargetFavoriteArticleError)
		return
	}

	// 5. 提交事务
	if err := tx.Commit().Error; err != nil {
		res.Fail(c, 500, consts.AffairCommitError)
		return
	}

	res.Success(c, nil, consts.MoveSuccess)
}

// RemoveFavoriteArticle 移除收藏夹博文
func (FavoriteApi) RemoveFavoriteArticle(c *gin.Context) {
	favoriteId := c.Param("id")
	var removeFavoriteArticleRequestParams request.RemoveFavoriteArticleRequestParams
	if err := c.ShouldBindJSON(&removeFavoriteArticleRequestParams); err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	err := db.Model(&models.FavoriteArticles{}).
		Where("favorite_id = ? And article_id in (?)", favoriteId, removeFavoriteArticleRequestParams.ArticleIDs).
		Delete(&models.FavoriteArticles{}).Error
	if err != nil {
		res.Fail(c, 500, consts.RemoveError)
	}
	res.Success(c, nil, consts.RemoveSuccess)
}

// DeleteFavorite 删除收藏夹
func (FavoriteApi) DeleteFavorite(c *gin.Context) {
	favoriteId := c.Param("id")
	db := global.MysqlDB
	tx := db.Begin()

	// 1. 删除收藏夹中的文章关联
	if err := tx.Where("favorite_id = ?", favoriteId).Delete(&models.FavoriteArticles{}).Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.DeleteFavoriteError)
		return
	}

	// 2. 删除收藏夹本身
	if err := tx.Where("id = ?", favoriteId).Delete(&models.Favorite{}).Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.DeleteFavoriteError)
		return
	}

	// 3. 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		res.Fail(c, 500, consts.DeleteFavoriteError)
		return
	}

	res.Success(c, nil, consts.DeleteFavoriteSuccess)
}
