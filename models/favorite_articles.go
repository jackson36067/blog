package models

// FavoriteArticles 收藏夹文章中间表
type FavoriteArticles struct {
	// 添加复合唯一索引
	FavoriteID uint `json:"favorite_id" gorm:"uniqueIndex:favorite_articles_key"`
	ArticleID  uint `json:"article_id" gorm:"uniqueIndex:favorite_articles_key"`
}
