package models

// Favorite 文章收藏夹表
type Favorite struct {
	Model
	Title    string    `json:"title" gorm:"size:32"`
	Abstract string    `json:"abstract" gorm:"size:255"` // 简介
	Coverage string    `json:"coverage" gorm:"size:255"`
	UserID   uint      `json:"user_id"`
	User     *User     `json:"-" gorm:"foreignKey:UserID;references:ID"`
	Articles []Article `json:"articles" gorm:"many2many:favorite_articles;joinForeignKey:favorite_id;joinReferences:article_id;"`
}
