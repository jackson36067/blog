package models

import "blog/enum"

// Article 文章表
type Article struct {
	Model
	Title         string             `json:"title" gorm:"size:32;"`
	Abstract      string             `json:"abstract" gorm:"size:255;"` // 文章简介
	Content       string             `json:"content" gorm:"type:text;"`
	CategoryID    uint               `json:"categoryId"` // 文章分类ID
	Category      *ArticleCategory   `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	TagList       []string           `json:"tagList" gorm:"type:longtext;serializer:json"` // 文章标签
	Coverage      string             `json:"coverage" gorm:"size:255;"`                    // 文章封面
	UserID        uint               `json:"userId"`
	User          *User              `json:"user" gorm:"foreignKey:UserID;references:ID"`                                                        // 与用户表为多对一关系
	BrowseCount   int                `json:"browserCount"`                                                                                       // 文章浏览量
	LikeCount     int                `json:"likeCount"`                                                                                          // 文章点赞数
	CommentCount  int                `json:"commentCount"`                                                                                       // 文章评论数
	CollectCount  int                `json:"collectCount"`                                                                                       // 文章收藏数
	PublicComment bool               `json:"publicComment"`                                                                                      // 是否开启评论
	Status        enum.ArticleStatus `json:"status"`                                                                                             // 文章状态: 草稿 审核中 已发布
	Favorites     []Favorite         `json:"favorites" gorm:"many2many:favorite_articles;joinForeignKey:article_id;joinReferences:favorite_id;"` // 文章被收藏收藏夹列表
	Comments      []Comment          `json:"comments" gorm:"foreignKey:ArticleID"`
}

// ArticleCategory 文章类型表
type ArticleCategory struct {
	ID       uint     `json:"id" gorm:"primary_key"`
	Title    string   `json:"title" gorm:"size:32;"`
	UserID   uint     `json:"userId"`
	User     *User    `json:"user" gorm:"foreignKey:UserID;references:ID"` // 与用户表为多对一关系
	Articles *Article `json:"articles" gorm:"foreignKey:CategoryID;references:ID"`
}
