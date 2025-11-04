package main

import (
	"blog/core"
	"blog/flags"
	"blog/global"
	"blog/routers"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	start := time.Now() // 记录启动开始时间
	// 解析命令行参数
	flags.ParseFlag()
	// 读取配置文件
	global.Conf = core.ReadConf()
	// 配置Logrus
	core.InitLogrus()
	// 初始化Mysql数据库连接
	global.MysqlDB = core.InitMysql()
	// 初始化redis连接
	global.RedisDB, global.Ctx = core.InitRedis()
	// 数据库迁移
	flags.Run()
	logrus.Infof("启动用时: %s", time.Since(start))
	routers.Run()
}
