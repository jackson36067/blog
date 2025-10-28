package models

// UserArticleBrowseHistory 用户浏览文章记录表
type UserArticleBrowseHistory struct {
	Model
	UserID    uint     `json:"user_id"`
	User      *User    `json:"user" gorm:"foreignKey:UserID"`
	ArticleID uint     `json:"article_id"`
	Article   *Article `json:"article" gorm:"foreignKey:ArticleID"`
}
