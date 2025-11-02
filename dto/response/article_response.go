package response

import "blog/models"

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
