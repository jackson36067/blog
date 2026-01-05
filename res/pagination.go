package res

import "gorm.io/gorm"

type Pagination struct {
	Page          int   `json:"page" form:"page"`
	PageSize      int   `json:"pageSize" form:"pageSize"`
	TotalElements int64 `json:"totalElements"` // 总共元素数量
	TotalPages    int   `json:"totalPages"`    // 总共页码数
	Data          any   `json:"data"`
}

func NewPagination(page int, pageSize int, totalElements int64, totalPages int, data any) *Pagination {
	return &Pagination{Page: page, PageSize: pageSize, TotalElements: totalElements, TotalPages: totalPages, Data: data}
}

// Paginate 通用分页器
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
