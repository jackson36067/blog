package models

// UserLogin 用户登录表
type UserLogin struct {
	Model
	UserID uint   `json:"userId"`
	User   *User  `json:"user" gorm:"foreignKey:UserID"`
	IP     string `json:"ip" gorm:"size:32"`
	Addr   string `json:"addr" gorm:"size:64"`
	UA     string `json:"ua" gorm:"size:128"` //使用什么端登录
}
