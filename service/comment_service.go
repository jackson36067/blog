package service

import (
	"blog/models"

	"gorm.io/gorm"
)

// GetAllChildIDs 获取某评论下所有子孙评论 ID（包含自身）
func GetAllChildIDs(db *gorm.DB, id uint) ([]uint, error) {
	ids := []uint{id}
	queue := []uint{id}

	for len(queue) > 0 {
		var childIDs []uint

		// 取队列第一个
		current := queue[0]
		queue = queue[1:]

		// 获取当前节点的直接子评论
		err := db.Model(&models.Comment{}).
			Where("parent_id = ?", current).
			Pluck("id", &childIDs).Error
		if err != nil {
			return nil, err
		}

		// 加入结果集和队列
		ids = append(ids, childIDs...)
		queue = append(queue, childIDs...)
	}

	return ids, nil
}
