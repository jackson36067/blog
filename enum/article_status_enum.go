package enum

type ArticleStatus int8

const (
	Draft       ArticleStatus = 1
	UnderReview ArticleStatus = 2
	Published   ArticleStatus = 3
)
