package models

import "time"

// UserArticleCollect 用户收藏文章表
type UserArticleCollect struct {
	// 创建复合索引一篇文章可以被用户收藏到多个收藏夹
	UserID     uint      `json:"userId" gorm:"uniqueIndex:idx_user_article_collect"`
	User       *User     `json:"user" gorm:"foreignKey:UserID"`
	ArticleID  uint      `json:"articleId" gorm:"uniqueIndex:idx_user_article_collect"`
	Article    *Article  `json:"article" gorm:"foreignKey:ArticleID"`
	FavoriteID uint      `json:"favoriteId" gorm:"uniqueIndex:idx_user_article_collect"` // 收藏夹ID
	Favorite   *Favorite `json:"-" gorm:"foreignKey:FavoriteID;"`
	CreatedAt  time.Time `json:"createdAt"`
}
