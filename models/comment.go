package models

// Comment 用户评论表
type Comment struct {
	Model
	Content        string     `json:"content" gorm:"type:text;"`
	UserID         uint       `json:"userId"`
	User           *User      `json:"user" gorm:"foreignKey:UserID"`
	ArticleID      uint       `json:"articleId"`
	Article        *Article   `json:"article" gorm:"foreignKey:ArticleID;references:ID"`
	ParentID       *uint      `json:"parentId"` // 父评论ID
	ParentComment  *Comment   `json:"-" gorm:"foreignKey:ParentID"`
	SubCommentList []*Comment `json:"-" gorm:"foreignKey:ParentID"`
	RootParentID   *uint      `json:"rootParentId"` // 根评论ID
	LikeCount      uint       `json:"likeCount"`    // 评论点赞数
}
