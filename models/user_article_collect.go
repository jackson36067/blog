package models

import "time"

// UserArticleCollect 用户收藏文章表
type UserArticleCollect struct {
	UserID     uint      `json:"userId" gorm:"uniqueIndex:idx_user_article_collect,priority:1"`
	User       *User     `json:"user" gorm:"foreignKey:UserID"`
	ArticleID  uint      `json:"articleId" gorm:"uniqueIndex:idx_user_article_collect,priority:2"`
	Article    *Article  `json:"article" gorm:"foreignKey:ArticleID"`
	FavoriteID uint      `json:"favoriteId" gorm:"uniqueIndex:idx_user_article_collect,priority:3"`
	Favorite   *Favorite `json:"-" gorm:"foreignKey:FavoriteID"`
	CreatedAt  time.Time `json:"createdAt"`
}
