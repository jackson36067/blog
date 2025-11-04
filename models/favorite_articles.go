package models

import "time"

// FavoriteArticles 收藏夹文章中间表
type FavoriteArticles struct {
	// 添加复合唯一索引
	FavoriteID uint      `json:"favoriteId" gorm:"uniqueIndex:favorite_articles_key"`
	ArticleID  uint      `json:"articleId" gorm:"uniqueIndex:favorite_articles_key"`
	CreatedAt  time.Time `json:"createdAt"`
}
