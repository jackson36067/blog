package service

import (
	"blog/dto/response"
	"blog/models"
	"blog/utils"
	"sort"
	"time"
)

// GroupArticlesByYearAndMonth 按年份、月份统计文章数量
// 返回结果包含每年的总数、每月的数量以及每月的时间范围
func GroupArticlesByYearAndMonth(articleList []response.ArticleStatistic) []response.ArticleYearStat {
	// 定义创作文历程 map
	// 第一层 key 为年份，第二层 key 为月份，value 为该月文章数量
	statsMap := make(map[int]map[int]int)

	for _, a := range articleList {
		// 获取文章的创建年份与月份
		year := a.CreatedAt.Year()
		month := int(a.CreatedAt.Month())

		// 如果该年份还没有初始化，先创建对应的月份 map
		if _, ok := statsMap[year]; !ok {
			statsMap[year] = make(map[int]int)
		}

		// 累加该年份该月份的文章数量
		statsMap[year][month]++
	}

	// 定义最终返回的结果切片
	var result []response.ArticleYearStat

	// 遍历年份 map，生成按年统计数据
	for year, monthMap := range statsMap {
		var months []response.ArticleMonthStat
		total := 0 // 当前年份的总文章数

		// 提取当前年份的所有月份 key（方便排序）
		monthKeys := make([]int, 0, len(monthMap))
		for m, _ := range monthMap {
			monthKeys = append(monthKeys, m)
		}
		// 对月份进行升序排序（1 月到 12 月）
		sort.Ints(monthKeys)

		// 遍历该年份的月份数据
		for _, m := range monthKeys {
			count := monthMap[m]
			total += count // 累加全年总数

			// 计算当前月的起止时间（方便后续前端查询用）
			start := time.Date(year, time.Month(m), 1, 0, 0, 0, 0, time.Local)
			// 下个月的第一天减 1 纳秒，即为当月的结束时间
			end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)

			// 构建月份统计信息
			months = append(months, response.ArticleMonthStat{
				Month:     m,
				Count:     count,
				StartTime: start,
				EndTime:   end,
			})
		}

		// 汇总当前年份的统计结果
		result = append(result, response.ArticleYearStat{
			Year:       year,
			TotalCount: total,
			Months:     months,
		})
	}

	// 将年份按降序排列（最近的年份排在最前）
	sort.Slice(result, func(i, j int) bool {
		return result[i].Year > result[j].Year
	})

	return result
}

// ArticlesToArticleResponse 将article切片数据转换成article-response切片数据
func ArticlesToArticleResponse(articles []models.Article) []response.ArticleResponse {
	return utils.MapSlice(articles, func(article models.Article) response.ArticleResponse {
		return response.ArticleResponse{
			Id:            article.ID,
			Title:         article.Title,
			Abstract:      article.Abstract,
			Content:       article.Content,
			Coverage:      article.Coverage,
			Tags:          article.TagList,
			CreatedAt:     article.CreatedAt.Format("2006-01-02 15:04:05"),
			BrowseCount:   article.BrowseCount,
			LikeCount:     article.LikeCount,
			CommentCount:  article.CommentCount,
			CollectCount:  article.CollectCount,
			PublicComment: article.PublicComment,
			Username:      article.User.Username,
			Avatar:        article.User.Avatar,
		}
	})
}

// GetArticleGroupedByTime 获取文章通过时间分组结果
func GetArticleGroupedByTime(browseArticles []models.UserArticleBrowseHistory) []response.ArticleGroup {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startOfYesterday := startOfToday.AddDate(0, 0, -1)
	startOfWeek := startOfToday.AddDate(0, 0, -7)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())

	groupMap := make(map[string][]response.ArticleResponse)
	orderKeys := make([]string, 0) // 👉 保存出现顺序

	for _, a := range browseArticles {
		// 根据浏览时间排序
		t := a.CreatedAt

		var groupKey string
		switch {
		case t.After(startOfToday):
			groupKey = "今日"
		case t.After(startOfYesterday):
			groupKey = "昨天"
		case t.After(startOfWeek):
			groupKey = "最近一周"
		case t.After(startOfYear):
			groupKey = t.Format("01-02") // 本年显示 MM-dd
		default:
			groupKey = t.Format("2006-01-02") // 往年显示 yyyy-MM-dd
		}
		article := a.Article
		ar := response.ArticleResponse{
			Id:            article.ID,
			Title:         article.Title,
			Abstract:      article.Abstract,
			Content:       article.Content,
			Coverage:      article.Coverage,
			Tags:          article.TagList,
			CreatedAt:     article.CreatedAt.Format("2006-01-02 15:04:05"),
			BrowseCount:   article.BrowseCount,
			LikeCount:     article.LikeCount,
			CommentCount:  article.CommentCount,
			CollectCount:  article.CollectCount,
			PublicComment: article.PublicComment,
			Username:      a.User.Username,
			Avatar:        a.User.Avatar,
		}
		// 首次遇到该分组时记录顺序
		if _, ok := groupMap[groupKey]; !ok {
			orderKeys = append(orderKeys, groupKey)
		}
		groupMap[groupKey] = append(groupMap[groupKey], ar)
	}

	// 构造返回结果，保持顺序（从新到旧）
	result := make([]response.ArticleGroup, 0, len(groupMap))
	// 按记录顺序构建结果（此顺序即为按 created_at 排序的顺序）
	for _, groupKey := range orderKeys {
		result = append(result, response.ArticleGroup{
			GroupTime: groupKey,
			Articles:  groupMap[groupKey],
		})
	}

	return result
}
