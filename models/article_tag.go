package models

type ArticleTag struct {
	Model
	Title       string `gorm:"size:32;not null" json:"title"`
	BrowseCount int    `json:"browseCount"` // 标签访问量
	PId         uint   `json:"pid"`         // 标签父ID
}
