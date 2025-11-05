package request

type UserRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Username string `form:"username"`
}

type UserFollowRequest struct {
	IsFollow bool `json:"isFollow"`
}
