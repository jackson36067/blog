package request

type UserRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Username string `form:"username"`
}

type UserFollowRequest struct {
	IsFollow bool `json:"isFollow"`
}

type UserCommentRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Username string `form:"username"`
	Type     string `form:"type"` // in: 收到的评论 out: 发出的评论
}
