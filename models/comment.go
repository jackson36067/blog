package models

// Comment 用户评论表
type Comment struct {
	Model
	Content       string   `json:"content" gorm:"type:text;charset:utf8mb4;collate:utf8mb4_0900_ai_ci"` // 满足可以存储表情
	UserID        uint     `json:"userId"`
	User          *User    `json:"user" gorm:"foreignKey:UserID"`
	ArticleID     uint     `json:"articleId"`
	Article       *Article `json:"article" gorm:"foreignKey:ArticleID;"`
	ParentID      *uint    `json:"parentId"` // 父评论ID
	ParentComment *Comment `json:"-" gorm:"foreignKey:ParentID"`
	RootParentID  *uint    `json:"rootParentId"` // 根评论ID
	LikeCount     uint     `json:"likeCount"`    // 评论点赞数
	// 新增字段：被回复的用户 ID
	ReplyToUserID *uint `json:"replyToUserId"`
	ReplyToUser   User  `gorm:"foreignKey:ReplyToUserID"`
}
