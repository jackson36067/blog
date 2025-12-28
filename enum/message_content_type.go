package enum

type MessageContentType int8

const (
	TextMessageContentType  MessageContentType = 1 // 文本消息
	ImageMessageContentType MessageContentType = 2 // 图片消息
	LinkMessageContentType  MessageContentType = 3 // 链接消息
	JsonMessageContentType  MessageContentType = 4 // JSON消息
)
