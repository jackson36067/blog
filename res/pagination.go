package res

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
