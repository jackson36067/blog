package models

type Banner struct {
	Model
	Coverage string `json:"coverage" gorm:"size:255"`
	Href     string `json:"href" gorm:"size:255"`
}
