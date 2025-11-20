package response

type UserDataResponse struct {
	OriginArticle               int      `json:"originArticle"`               // 原创文章数量
	Fans                        int      `json:"fans"`                        // 粉丝数量
	Follow                      int      `json:"follow"`                      // 关注数
	IP                          string   `json:"ip"`                          // IP地址
	JoinTime                    string   `json:"joinTime"`                    // 加入博客时间
	CodeAge                     int      `json:"codeAge"`                     // 码龄
	Username                    string   `json:"username"`                    // 用户名
	Avatar                      string   `json:"avatar"`                      // 头像
	Sex                         int8     `json:"sex"`                         // 性别
	Abstract                    string   `json:"abstract"`                    // 简介
	Birthday                    string   `json:"birthday"`                    // 出生日期
	HobbyTags                   []string `json:"hobbyTags"`                   // 兴趣标签
	PublicFanList               bool     `json:"publicFanList"`               // 公开粉丝列表
	PublicCollectList           bool     `json:"publicCollectList"`           // 公开收藏列表
	PublicFollowList            bool     `json:"publicFollowList"`            // 公开关注列表
	SinceLastUpdateUsernameDays int      `json:"sinceLastUpdateUsernameDays"` // 距上一次更新用户名天数
}

type UserAchievementResponse struct {
	TotalLikes    int64 `json:"totalLikes"`
	TotalCollects int64 `json:"totalCollects"`
	TotalComments int64 `json:"totalComments"`
}

type UserFollowResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Abstract string `json:"abstract"`
	IsFollow bool   `json:"isFollow"`
}

type UserCommentResponse struct {
	CommentID uint   `json:"commentId"` // 评论id
	ArticleID uint   `json:"articleId"` // 文章id
	Content   string `json:"content"`   // 文章内容
	Title     string `json:"title"`     // 评论文章标题
	CreatedAt string `json:"createdAt"` // 评论时间
}

type UserLoginLogResponse struct {
	ID        uint   `json:"id"`
	Ip        string `json:"ip"`
	Addr      string `json:"addr"`
	CreatedAt string `json:"createdAt"`
}
