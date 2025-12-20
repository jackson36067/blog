package request

import "blog/enum"

type MessageChatHistoryRequestParams struct {
	ChatUserId uint `form:"chatUserId"`
}

type UpdateSessionRequestParams struct {
	IsMuted  *bool `json:"isMuted"`  // 是否静音
	IsPinned *bool `json:"isPinned"` // 是否置顶
}

type OtherMessageRequestParams struct {
	Page     int              `form:"page"`
	PageSize int              `form:"pageSize"`
	Type     enum.MessageType `form:"type"` // 2评论@ 3点赞收藏 4粉丝
}
