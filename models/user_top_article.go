package models

import "time"

// UserTopArticle 用户置顶文章表
type UserTopArticle struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"userId" gorm:"uniqueIndex:idx_user_top_article"`
	User      *User     `json:"-" gorm:"foreignKey:UserID;references:ID"`
	ArticleID uint      `json:"articleId" gorm:"uniqueIndex:idx_user_top_article"`
	Article   *Article  `json:"-" gorm:"foreignKey:ArticleID;references:ID"`
	CreatedAt time.Time `json:"createdAt"`
}
