package response

import (
	"blog/models"
	"time"
)

type ArticleResponse struct {
	Id            uint     `json:"id"` // 文章id
	Title         string   `json:"title"`
	Abstract      string   `json:"abstract"`
	Content       string   `json:"content"`
	Coverage      string   `json:"coverage"`
	Tags          []string `json:"tags"`
	CreatedAt     string   `json:"createdAt"`
	BrowseCount   int      `json:"browseCount"`
	LikeCount     int      `json:"likeCount"`
	CommentCount  int      `json:"commentCount"`
	CollectCount  int      `json:"collectCount"`
	PublicComment bool     `json:"publicComment"`
}

type ArticleHotTagsAndRandCategoryResponse struct {
	ArticleTags       []models.ArticleTag      `json:"articleTags"`
	ArticleCategories []models.ArticleCategory `json:"articleCategories"`
}

type ArticleStatistic struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type ArticleMonthStat struct {
	Month     int       `json:"month"`
	Count     int       `json:"count"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type ArticleYearStat struct {
	Year       int                `json:"year"`
	TotalCount int                `json:"totalCount"`
	Months     []ArticleMonthStat `json:"months"`
}

type ArticleGroup struct {
	GroupTime string            `json:"groupTime"`
	Articles  []ArticleResponse `json:"articles"`
}
