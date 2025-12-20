package response

type GroupedChatResponse struct {
	SendTime string         `json:"sendTime"`
	Messages []ChatResponse `json:"messages"`
}

type SessionResponse struct {
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

type ChatResponse struct {
	ID             uint   `json:"id"`             // 消息ID
	UserID         uint   `json:"userId"`         // 用户id
	Username       string `json:"username"`       // 用户名称
	UserAvatar     string `json:"userAvatar"`     // 用户头像
	ChatUserID     uint   `json:"chatUserId"`     // 聊天用户id
	ChatUsername   string `json:"chatUsername"`   // 聊天用户名称
	ChatUserAvatar string `json:"chatUserAvatar"` // 聊天用户头像
	Message        string `json:"message"`        // 消息内容
}

type OtherMessageResponse struct {
	ID            uint   `json:"id"`            // 消息ID
	UserID        uint   `json:"userId"`        // 发送消息的用户id
	Username      string `json:"username"`      // 发送消息的用户名称
	UserAvatar    string `json:"userAvatar"`    // 发送消息的用户头像
	Message       string `json:"message"`       // 消息内容
	ActionMessage string `json:"actionMessage"` // 动作消息（如点赞、收藏、系统消息等）
	SendTime      string `json:"sendTime"`      // 发送时间
	Extra         string `json:"extra"`         // 扩展字段（如文章ID、链接、评论ID等）
	Title         string `json:"title"`         // 标题(可以是文章标题、评论标题等)
	IsFollow      bool   `json:"isFollow"`      // 是否关注发送消息的用户
}
