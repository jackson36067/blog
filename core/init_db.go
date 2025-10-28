package core

import (
	"blog/global"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitMysql 初始化数据库连接
func InitMysql() (db *gorm.DB) {
	mysqlConnector := global.Conf.DB.Mysql
	db, err := gorm.Open(mysql.Open(mysqlConnector.MysqlDSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			// 生成表的时候表明不带上复数
			SingularTable: true,
			// 字段名小写+下划线
			NameReplacer: nil,
			NoLowerCase:  false,
		},
	})
	if err != nil {
		logrus.Fatalf("mysql数据库连接失败 %s", err.Error())
	}
	// 设置连接池大小
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("mysql数据库连接成功")
	return
}

func InitRedis() (*redis.Client, context.Context) {
	redisConf := global.Conf.DB.Redis
	ctx := context.Background()
	Rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + redisConf.Port,
		Password: redisConf.Password,
		DB:       redisConf.Database,
		PoolSize: 10, // 连接池大小
	})
	// 测试连接
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Fatalf("redis连接失败%s", err.Error())
	}
	logrus.Infof("redis连接成功")
	return Rdb, ctx
}
