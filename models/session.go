package models

import "time"

type Session struct {
	Model
	UserIDA        uint       `json:"userIdA"` // 用户 A
	UserA          User       `json:"-" gorm:"foreignKey:UserIDA"`
	UserIDB        uint       `json:"userIdB"` // 用户 B
	UserB          User       `json:"-" gorm:"foreignKey:UserIDB"`
	LatestChatTime time.Time  `json:"latestChatTime"` // 最新时间
	LatestMessage  string     `json:"latestMessage"`  // 最新消息
	IsPinnedA      bool       `json:"isPinnedA"`      // A 是否置顶
	IsPinnedB      bool       `json:"isPinnedB"`      // B 是否置顶
	IsMutedA       bool       `json:"isMutedA"`       // A 是否免打扰
	IsMutedB       bool       `json:"isMutedB"`       // B 是否免打扰
	DeletedAtA     *time.Time `json:"deletedAtA"`     // 用户a删除会话时间
	DeletedAtB     *time.Time `json:"deletedAtB"`     // 用户b删除会话时间
}
