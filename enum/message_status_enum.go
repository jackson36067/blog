package enum

type MessageStatus int8

const (
	Normal    MessageStatus = 0 // 正常
	Withdraw  MessageStatus = 1 // 撤回
	Failure   MessageStatus = 2 // 失效
	PushError MessageStatus = 3 // 推送失败
)
