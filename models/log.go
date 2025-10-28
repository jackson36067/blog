package models

import "blog/enum"

// Log 日志表
type Log struct {
	Model
	Type    enum.LogType  `json:"type"` // 日志类型
	Title   string        `json:"title" gorm:"size:64"`
	Content string        `json:"content" gorm:"type:text;"`
	Level   enum.LogLevel `json:"level"`  // 日志级别
	UserID  uint          `json:"userId"` // 用户id
	User    *User         `json:"-" gorm:"foreignKey:UserID"`
	IP      string        `json:"ip" gorm:"size:32"`
	Addr    string        `json:"addr" gorm:"size:64"`
	IsRead  bool          `json:"isRead"` // 是否已读
	// 只有日志类行为登录类型时才有
	LoginStatus bool           `json:"loginStatus"`              // 登录状态
	Username    string         `json:"username" gorm:"size:32"`  // 登录用户名
	Password    string         `json:"password" gorm:"size:255"` // 登录密码
	LoginType   enum.LoginType `json:"loginType"`                // 登录类型
}
