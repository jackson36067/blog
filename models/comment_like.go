package models

import "time"

// CommentLike 评论点赞表
type CommentLike struct {
	// 创建复合索引uniqueIndex:idx_comment_likes
	UserID    uint      `json:"userId" gorm:"uniqueIndex:idx_comment_likes"`
	User      *User     `json:"user" gorm:"foreignKey:UserID"`
	CommentID uint      `json:"commentId" gorm:"uniqueIndex:idx_comment_likes"`
	Comment   *Comment  `json:"comment" gorm:"foreignKey:CommentID"`
	CreatedAt time.Time `json:"createdAt"`
}
