package flags

import "flag"

type Options struct {
	// 加载哪个配置文件
	ConfigFile string
	// 是否迁移数据库
	MigrateDB bool
}

var FlagOptions = new(Options)

// ParseFlag 读取命令行配置
func ParseFlag() {
	flag.StringVar(&FlagOptions.ConfigFile, "f", "settings.yml", "配置文件")
	flag.BoolVar(&FlagOptions.MigrateDB, "db", false, "迁移数据库")
	flag.Parse()
}
