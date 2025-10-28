package main

import (
	"blog/core"
	"blog/flags"
	"blog/global"
	"blog/routers"
)

func main() {
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
	routers.Run()
}
