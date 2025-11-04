package response

type FavoriteListResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Abstract     string `json:"abstract"`
	IsDefault    bool   `json:"isDefault"`
	ArticleCount int    `json:"articleCount"`
}

type FavoriteArticleResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
