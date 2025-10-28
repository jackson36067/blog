package models

import "time"

// ArticleLike 文章点赞表
type ArticleLike struct {
	// 创建复合索引uniqueIndex:idx_article_likes
	UserID    uint      `json:"userId" gorm:"uniqueIndex:idx_article_likes"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
	ArticleID uint      `json:"articleId" gorm:"uniqueIndex:idx_article_likes"`
	Article   *Article  `json:"article" gorm:"foreignKey:ArticleID;references:ID"`
	CreatedAt time.Time `json:"createdAt"`
}
