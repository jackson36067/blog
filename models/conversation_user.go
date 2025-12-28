package models

import "time"

type ConversationUser struct {
	Model
	ConversationID uint         `json:"conversationId" gorm:"not null;uniqueIndex:idx_conversation_user"` // 会话 id
	Conversation   Conversation `json:"-" gorm:"foreignKey:ConversationID"`
	UserID         uint         `json:"userId" gorm:"not null;uniqueIndex:idx_conversation_user"` // 用户 id
	User           User         `json:"-" gorm:"foreignKey:UserID"`
	IsPinned       bool         `json:"isPinned"`  // 是否置顶
	IsMuted        bool         `json:"isMuted"`   // 是否免打扰
	DeletedAt      *time.Time   `json:"deletedAt"` // 删除时间
}
