package consts

import "time"

// 用户相关的常量
const (
	UsernameOrPwdNull = "用户名或密码不能为空"
	UserNotFound      = "用户不存在"
	PwdError          = "密码错误"
	UsernameExist     = "用户名已被使用, 请更换一个"
)

// 登录相关的常量
const (
	LoginSuccess    = "登录成功"
	UnknowLoginType = "未知的登录方式"
)

// 注册相关常量
const (
	RegisterSuccess = "注册成功"
)

// Redis相关常量
const (
	// EmailCodeKeyPrefix 以邮箱为结尾
	EmailCodeKeyPrefix = "email:code:"
)

// Jwt相关常量
const (
	Issuer         = "gin-jwt-demo"
	ExpireDuration = time.Hour * 24 * 7
	Subject        = "user token"
	InValidToken   = "无效token"
	GenJwtError    = "生成jwt令牌失败"
)

// 邮箱相关常量
const (
	SendMailError   = "发送邮件失败"
	SendMailSuccess = "邮件发送成功"
	EmailOrCodeNull = "邮箱或验证码不能为空"
	EmailExist      = "邮箱已被注册, 请更换"
	CodeError       = "验证码错误, 请重新获取"
)
