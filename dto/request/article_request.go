package request

import (
	"blog/enum"
	"time"
)

type ArticleQueryParams struct {
	Page       int      `form:"page" binding:"required"`
	PageSize   int      `form:"pageSize" binding:"required"`
	Title      string   `form:"title"`      // 文章标题
	CategoryId uint     `form:"categoryId"` // 文章分类
	Tags       []string `form:"tags"`       // 文章标签
	UserId     uint     `form:"userId"`     // 文章所属用户
}

type MyArticleQueryParams struct {
	Page       int                    `form:"page" binding:"required"`
	PageSize   int                    `form:"pageSize" binding:"required"`
	Visibility enum.ArticleVisibility `form:"visibility"` // 获取文章的可见范围
	OrderBy    string                 `form:"orderBy"`    // 根据什么排序
	OrderType  string                 `form:"orderType"`  // 排序方式 (升序,降序)
	StartTime  time.Time              `form:"startTime"`
	EndTime    time.Time              `form:"endTime"`
}
