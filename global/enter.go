package global

import (
	"blog/conf"
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Conf    *conf.Conf
	MysqlDB *gorm.DB
	RedisDB *redis.Client
	Ctx     = context.Background()
)
