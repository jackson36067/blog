package utils

import "time"

func GetDetailDate(t time.Time) string {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startOfYesterday := startOfToday.AddDate(0, 0, -1)
	startOfWeek := startOfToday.AddDate(0, 0, -7)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	var detailDate string
	switch {
	case t.After(startOfToday):
		detailDate = "今日"
	case t.After(startOfYesterday):
		detailDate = "昨天"
	case t.After(startOfWeek):
		detailDate = "最近一周"
	case t.After(startOfYear):
		detailDate = t.Format("01-02") // 本年显示 MM-dd
	default:
		detailDate = t.Format("2006-01-02") // 往年显示 yyyy-MM-dd
	}
	return detailDate
}

func FormatChatTime(t time.Time) string {
	now := time.Now()

	// 今日 00:00
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startOfYesterday := startOfToday.AddDate(0, 0, -1)
	startOfTwoDaysAgo := startOfToday.AddDate(0, 0, -2)
	startOfThreeDaysAgo := startOfToday.AddDate(0, 0, -3)

	// 今天
	if t.After(startOfToday) {
		return t.Format("15:04")
	}

	// 昨天
	if t.After(startOfYesterday) {
		return "昨天"
	}

	// 前天（2 天前）
	if t.After(startOfTwoDaysAgo) {
		return "前天"
	}

	// 大前天（3 天前）
	if t.After(startOfThreeDaysAgo) {
		return "三天前"
	}

	// 超过三天 —— 今年
	if t.Year() == now.Year() {
		return t.Format("01-02")
	}

	// 不是今年 —— 显示完整年月日
	return t.Format("2006-01-02")
}
