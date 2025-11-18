package request

type UserRequestParams struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Username string `form:"username"`
}

type UserFollowRequestParam struct {
	IsFollow bool `json:"isFollow"`
}

type UserCommentRequestParams struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Username string `form:"username"`
	Type     string `form:"type"` // in: 收到的评论 out: 发出的评论
}

type UpdateUserRequestParams struct {
	UserID            uint     `json:"userId"`
	Avatar            string   `json:"avatar"`
	Username          string   `json:"username"`
	Sex               *int8    `json:"sex"`
	Abstract          string   `json:"abstract"`
	Birthday          string   `json:"birthday"`
	HobbyTags         []string `json:"hobbyTags"`
	PublicFanList     *bool    `json:"publicFanList"`
	PublicCollectList *bool    `json:"publicCollectList"`
	PublicFollowList  *bool    `json:"publicFollowList"`
	Email             string   `json:"email"`
	Password          string   `json:"password"`
}
