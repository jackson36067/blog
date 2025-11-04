package request

type NewFavoriteRequestParams struct {
	Title     string `json:"title"`
	Abstract  string `json:"abstract"`
	IsDefault bool   `json:"isDefault"`
}

type UpdateFavoriteRequestParams struct {
	Title     string `json:"title"`
	Abstract  string `json:"abstract"`
	IsDefault bool   `json:"isDefault"`
}

type MoveFavoriteRequestParams struct {
	SourceFavoriteID uint   `json:"sourceFavoriteId"` // 来源收藏夹ID
	TargetFavoriteID uint   `json:"targetFavoriteId"` // 目标收藏夹ID
	ArticleIDs       []uint `json:"articleIds"`       // 要移动的文章ID
}

type RemoveFavoriteArticleRequestParams struct {
	ArticleIDs []uint `json:"articleIds"`
}
