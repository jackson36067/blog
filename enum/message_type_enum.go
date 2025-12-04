package enum

type MessageType int8

const (
	PrivateMessage       MessageType = 1 // 私信
	CommentMessage       MessageType = 2 // 评论@
	LikeOrCollectMessage MessageType = 3 // 点赞收藏
	FanMessage           MessageType = 4 // 粉丝
	SystemMessage        MessageType = 5 // 系统
)
