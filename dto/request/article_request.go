package request

type ArticleQueryParams struct {
	Page       int      `form:"page" binding:"required"`
	PageSize   int      `form:"pageSize" binding:"required"`
	Title      string   `form:"title"`      // 文章标题
	CategoryId uint     `form:"categoryId"` // 文章分类
	Tags       []string `form:"tags"`       // 文章标签
	UserId     uint     `form:"userId"`     // 文章所属用户
}
