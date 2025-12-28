package models

import (
	"time"

	"blog/enum"
)

type Message struct {
	Model
	Type          enum.MessageType        `json:"type"`       // 消息类型：1私信 2评论@ 3点赞收藏 4粉丝 5系统
	SendUserID    uint                    `json:"sendUserId"` // 发送者ID
	SendUser      User                    `json:"-" gorm:"foreignKey:SendUserID"`
	ReceiveUserID uint                    `json:"receiveUserId"` // 接收者ID
	ReceiveUser   User                    `json:"-" gorm:"foreignKey:ReceiveUserID"`
	Content       string                  `json:"content" gorm:"type:text;charset:utf8mb4;collate:utf8mb4_0900_ai_ci"` // 消息内容（文本或JSON）
	ContentType   enum.MessageContentType `json:"contentType"`                                                         // 消息内容类型：1文本 2图片 3链接 4JSON
	Extra         string                  `json:"extra"`                                                               // 扩展字段（如文章ID、链接、评论ID等）
	ActionMessage string                  `json:"actionMessage"`                                                       // 动作消息（如点赞、收藏、系统消息等）
	IsRead        bool                    `json:"isRead"`                                                              // 是否已读
	ReadTime      *time.Time              `json:"readTime"`                                                            // 阅读时间
	SendTime      time.Time               `json:"sendTime"`                                                            // 发送时间
	PlanPushTime  *time.Time              `json:"planPushTime"`                                                        // 计划推送时间（延迟推送）
	RealPushTime  *time.Time              `json:"realPushTime"`                                                        // 实际推送时间（MQ/WebSocket实际推送）
	Status        enum.MessageStatus      `json:"status"`                                                              // 状态：0正常 1撤回 2失效 3推送失败
	SessionID     *uint                   `json:"sessionId"`                                                           // 会话ID 只有当消息类型为私信时才有
	Session       *Session                `json:"-" gorm:"foreignKey:SessionID"`
}
