package global

import (
	"blog/conf"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	Conf    *conf.Conf
	MysqlDB *gorm.DB
	RedisDB *redis.Client
	Ctx     = context.Background()
	MQ      *amqp.Connection
)
