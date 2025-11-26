package request

type PublishCommentRequestParams struct {
	Content          string `json:"content" binding:"required"`          // 评论内容
	CommentArticleId uint   `json:"commentArticleId" binding:"required"` // 评论的文章id
	ParentCommentId  *uint  `json:"parentCommentId"`                     // 评论的父评论id (为根评论不存在)
	RootCommentId    *uint  `json:"rootCommentId"`                       // 评论的根评论id (为根评论不存在)
	ReplyToUserId    *uint  `json:"replyToUserId"`                       // 回复用户id (为根评论不存在)
}

type LikeCommentRequestParams struct {
	IsLike bool `json:"isLike"`
}

type DeleteCommentRequestParams struct {
	IsRoot bool `json:"isRoot"` // 是否为根目录
}
