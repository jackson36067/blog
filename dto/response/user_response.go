package response

type UserDataResponse struct {
	OriginArticle int    `json:"originArticle"` // 原创文章数量
	Fans          int    `json:"fans"`          // 粉丝数量
	Follow        int    `json:"follow"`        // 关注数
	IP            string `json:"ip"`            // IP地址
	JoinTime      string `json:"joinTime"`      // 加入博客时间
	CodeAge       int    `json:"codeAge"`       // 码龄
	Avatar        string `json:"avatar"`        // 头像
}

type UserAchievementResponse struct {
	TotalLikes    int64 `json:"totalLikes"`
	TotalCollects int64 `json:"totalCollects"`
	TotalComments int64 `json:"totalComments"`
}
