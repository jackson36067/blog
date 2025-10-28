package models

import "fmt"

// Image 用户上传图片表
type Image struct {
	Model
	Filename string `json:"filename" gorm:"size:64"` // 文件名称
	Path     string `json:"path" gorm:"size:255"`    // 文件路径
	Size     int64  `json:"size"`                    // 文件大小
	Hash     string `json:"hash" gorm:"size:32"`
}

func (i Image) WebPath() string {
	return fmt.Sprintf("/")
}
