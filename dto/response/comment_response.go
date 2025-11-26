package response

type CommentResponse struct {
	ID              uint               `json:"id"`
	UserID          uint               `json:"userId"`
	Avatar          string             `json:"avatar"`
	Username        string             `json:"username"`
	Content         string             `json:"content"`
	LikeCount       uint               `json:"likeCount"`
	IsLike          bool               `json:"isLike"` // 是否点赞评论
	RootCommentID   *uint              `json:"rootCommentId"`
	SubComment      []*CommentResponse `json:"subComment"`
	CreateTime      string             `json:"createTime"`
	ReplyToUsername string             `json:"replyToUsername"`
}
