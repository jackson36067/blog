package request

type UserLikesRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
