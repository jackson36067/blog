package consts

import "time"

// 用户相关的常量
const (
	UsernameOrPwdNull = "用户名或密码不能为空"
	UserNotFound      = "用户不存在"
	PwdError          = "密码错误"
	UsernameExist     = "用户名已被使用, 请更换一个"
	UserId            = "userId"
	UnKnowUserIdType  = "无法识别的用户ID类型"
)

// 登录,注册相关的常量
const (
	LoginSuccess    = "登录成功"
	UnknowLoginType = "未知的登录方式"
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
	InvalidToken   = "无效token"
	GenJwtError    = "生成jwt令牌失败"
	TokenNull      = "token为空"
	ExpireToken    = "token已过期，请重新登录"
)

// 邮箱相关常量
const (
	SendMailError    = "发送邮件失败"
	SendMailSuccess  = "邮件发送成功"
	EmailOrCodeNull  = "邮箱或验证码不能为空"
	EmailExist       = "邮箱已被注册, 请更换"
	CodeError        = "验证码错误, 请重新获取"
	EmailSubject     = "jackson-blog登录验证码"
	SendEmailSuccess = "获取验证码成功"
)

// IP相关常量
const (
	IpParseError        = "获取Ip失败"
	InvalidIp           = "无效的IP地址"
	LoadIpDatabaseError = "ip地址数据库加载失败"
	Localhost           = "本机"
)

// RequestParamParseError 用于请求参数获取失败提示
const (
	RequestParamParseError = "请求参数解析失败"
)

// 响应信息相关常量
const (
	ClearUserBrowseHistorySuccess    = "清空历史成功"
	NewFavoriteSuccess               = "新建收藏夹成功"
	NoUpdateField                    = "未提供任何更新字段"
	UpdateSuccess                    = "更新成功"
	DeleteTargetFavoriteArticleError = "删除来源收藏夹记录失败"
	InsertNewRecordError             = "插入新记录失败"
	FindTargetFavoriteError          = "查询目标收藏夹失败"
	MoveSuccess                      = "移动成功"
	RemoveError                      = "移除博文失败"
	RemoveSuccess                    = "移除成功"
	DeleteFavoriteError              = "删除收藏夹失败"
	DeleteFavoriteSuccess            = "删除收藏夹成功"
)

// AffairCommitError 事务相关常量
const (
	AffairCommitError = "事务提交失败"
)
