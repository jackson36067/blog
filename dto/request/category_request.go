package request

type CategoryArticleRequestParam struct {
	Title    string `form:"title"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}
