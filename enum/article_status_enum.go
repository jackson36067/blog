package enum

type ArticleStatus int8

const (
	Draft       ArticleStatus = 1 // 草稿
	UnderReview ArticleStatus = 2 // 审查中
	Published   ArticleStatus = 3 // 已发布
)
