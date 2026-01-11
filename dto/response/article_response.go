package response

import (
	"time"

	"blog/enum"
	"blog/models"
)

type ArticleResponse struct {
	Id               uint                   `json:"id"` // 文章id
	Title            string                 `json:"title"`
	Abstract         string                 `json:"abstract"`
	Content          string                 `json:"content"`
	Coverage         string                 `json:"coverage"`
	Tags             []string               `json:"tags"`
	CreatedAt        string                 `json:"createdAt"`
	BrowseCount      int                    `json:"browseCount"`
	LikeCount        int                    `json:"likeCount"`
	CommentCount     int                    `json:"commentCount"`
	CollectCount     int                    `json:"collectCount"`
	PublicComment    bool                   `json:"publicComment"`
	UserID           uint                   `json:"userId"`
	Username         string                 `json:"username"`
	Avatar           string                 `json:"avatar"`
	IsLike           bool                   `json:"isLike"`    // 当前用户是否点赞文章
	IsCollect        bool                   `json:"isCollect"` // 当前用户是否收藏该文章
	IsFollow         bool                   `json:"isFollow"`  // 当前用户是否关注作者
	Comment          *CommentResponse       `json:"comment"`
	TotalComment     uint                   `json:"totalComment"`     // 文章总共评论数
	TotalRootComment uint                   `json:"totalRootComment"` // 文章总共根评论数
	CategoryName     string                 `json:"categoryName"`     // 文章分类名称
	Visibility       enum.ArticleVisibility `json:"visibility"`       // 文章可见范围
	Status           enum.ArticleStatus     `json:"status"`           // 文章状态
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

type ArticleTagTree struct {
	Id          uint              `json:"id"`
	Title       string            `json:"title"`
	BrowseCount int               `json:"browseCount"`
	Children    []*ArticleTagTree `json:"children"`
}
