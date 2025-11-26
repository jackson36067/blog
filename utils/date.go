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
