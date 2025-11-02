package enum

type ArticleVisibility int8

const (
	Public  ArticleVisibility = 0 // 全部可见
	Fans    ArticleVisibility = 1 // 粉丝可见
	Private ArticleVisibility = 2 // 仅我可见
)
