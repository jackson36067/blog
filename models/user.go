package models

import "time"

// User 用户表
type User struct {
	Model
	Username       string      `json:"username" gorm:"size:32;unique;not null"`
	Nickname       string      `json:"nickname" gorm:"size:32;not null"`
	Avatar         string      `json:"avatar" gorm:"size:255;"`
	Abstract       string      `json:"abstract" gorm:"size:255;"` // 简介
	Sex            int8        `json:"sex"`                       // 性别 0.male 1.female
	Birthday       time.Time   `json:"birthday"`                  // 生日
	RegisterSource int8        `json:"registerSource"`            // 注册来源
	Password       string      `json:"password" gorm:"size:120;"`
	CodeAge        int         `json:"codeAge" gorm:"size:255;"` // 码龄
	Email          string      `json:"email" gorm:"size:255;"`
	OpenId         string      `json:"openId" gorm:"size:64;"` // 第三方登录的唯一ip
	UserConfig     *UserConfig `json:"-" gorm:"foreignKey:UserID"`
	Articles       []Article   `json:"articles" gorm:"foreignKey:UserID"` // 用户文章列表
	Role           int8        `json:"role"`                              // 角色 1.管理员 2.普通用户 3.访客
}

// UserConfig 用户配置表
type UserConfig struct {
	ID                 uint      `json:"id" gorm:"primary_key"`
	UserID             uint      `json:"userId" gorm:"unique;not null"`
	HobbyTags          []string  `json:"hobbyTags" gorm:"type:longtext;serializer:json"` // 兴趣标签
	UpdateUsernameDate time.Time `json:"updateUsernameDate"`                             // 上次修改用户名时间
	PublicCollectList  bool      `json:"publicCollectList"`                              // 公开我的收藏列表
	PublicFollowList   bool      `json:"publicFollowList"`                               // 公开我的关注列表
	PublicFanList      bool      `json:"publicFanList"`                                  // 公开我的粉丝列表
	HomeStyleID        uint      `json:"homeStyleId"`                                    // 主页样式Id
}
