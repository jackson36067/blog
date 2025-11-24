package response

type CommentResponse struct {
	ID              uint               `json:"id"`
	UserID          uint               `json:"userId"`
	Avatar          string             `json:"avatar"`
	Username        string             `json:"username"`
	Content         string             `json:"content"`
	LikeCount       uint               `json:"likeCount"`
	RootCommentID   *uint              `json:"rootCommentId"`
	SubComment      []*CommentResponse `json:"subComment"`
	CreateTime      string             `json:"createTime"`
	ReplyToUsername string             `json:"replyToUsername"`
}
