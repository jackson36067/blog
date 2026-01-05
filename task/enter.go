package task

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// StartCronJobs 定义所有定时任务
func StartCronJobs(db *gorm.DB, rdb *redis.Client) {
	// 保证Redis缓存了key
	go SyncHighQualityArticlesWithDecay(db, rdb)

	c := cron.New(cron.WithSeconds())

	// 添加定时任务：每 10 分钟刷新一次
	_, err := c.AddFunc("0 */10 * * * *", func() {
		// 同步高质量文章引入时间衰减方案, 让内容保鲜
		SyncHighQualityArticlesWithDecay(db, rdb)
	})
	if err != nil {
		log.Fatal("定时任务添加失败:", err)
	}

	c.Start()
	logrus.Infof("定时任务启动成功")
}
