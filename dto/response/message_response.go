package response

type ChatResponse struct {
	SessionID      uint   `json:"sessionId"`      // 聊天队列ID
	ChatUserID     uint   `json:"chatUserId"`     // 聊天用户id
	ChatUsername   string `json:"chatUsername"`   // 聊天用户名称
	ChatUserAvatar string `json:"chatUserAvatar"` // 聊天用户头像
	IsPinned       bool   `json:"isPinned"`       // 是否置顶聊天
	IsMuted        bool   `json:"isMuted"`        // 是否静音聊天
	LatestMessage  string `json:"latestMessage"`  // 最近一条消息内容
	LatestChatTime string `json:"latestChatTime"` // 最近一条消息时间
	IsFollow       bool   `json:"isFollow"`       // 是否关注聊天用户
	UnreadCount    uint   `json:"unreadCount"`    // 未读消息数量
}
