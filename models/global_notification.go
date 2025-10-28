package models

// GlobalNotification 全局通知表
type GlobalNotification struct {
	Model
	Title   string `json:"title" gorm:"size:64"`
	Icon    string `json:"icon" gorm:"size:255"` // 图标
	Content string `json:"content" gorm:"type:text"`
	Href    string `json:"href" gorm:"size:255"` // 用户点击消息, 然后跳转的地址
}
