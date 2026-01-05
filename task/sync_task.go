package task

import (
	"fmt"
	"log"

	"blog/consts"
	"blog/global"
	"blog/models"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// SyncHighQualityArticlesWithDecay 同步高质量文章引入时间衰减方案, 让内容保鲜
func SyncHighQualityArticlesWithDecay(db *gorm.DB, rdb *redis.Client) {
	var results []struct {
		ID    uint    // 文章id
		Score float64 // 文章分数(依据收藏数, 点赞数, 收藏数, 时间衰减算法)
	}
	// G (Gravity) 重力因子，通常取 1.5 到 1.8。值越大，旧文章分数掉得越快。
	const G = 1.8
	// SQL 说明：
	// 1. 计算基础分 (Views + Likes*5 + Comments*10)
	// 2. TIMESTAMPDIFF 计算文章发布到现在经过的小时数
	// 3. Score = BaseScore / (Hours + 2)^G
	err := db.Model(&models.Article{}).
		Select(fmt.Sprintf(
			"id, (comment_count * 1 + like_count * 5 + collect_count * 10) / POW(TIMESTAMPDIFF(HOUR, created_at, NOW()) + 2, %f) as score",
			G,
		)). // 计算文章分数
		Order("score DESC").
		Limit(100).
		Scan(&results).Error
	if err != nil {
		log.Printf("同步失败: %v", err)
		return
	}
	// 如果获取到的文章结果为0, 那么不需要同步到Redis中
	if len(results) == 0 {
		return
	}
	// 使用Redis Pipeline批量更新
	pipe := rdb.Pipeline()
	ctx := global.Ctx

	// 先清空Redis中缓存的高质量文章
	pipe.Del(ctx, consts.HighQualityArticleKey)

	// 存储高质量文章结果到Redis中
	var members []redis.Z
	for _, item := range results {
		members = append(members, redis.Z{Score: item.Score, Member: item.ID})
	}
	pipe.ZAdd(ctx, consts.HighQualityArticleKey, members...)
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Printf("同步高质量文章失败: %v", err)
		return
	} else {
		logrus.Infof("同步高质量文章成功")
	}
}
