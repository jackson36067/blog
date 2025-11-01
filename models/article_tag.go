package models

type ArticleTag struct {
	Model
	Title       string `gorm:"size:32;not null;unique" json:"title"`
	BrowseCount int    `json:"browse_count"` // 标签访问量
}
