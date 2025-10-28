package core

import (
	"blog/conf"
	"blog/flags"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadConf 读取配置文件中的配置到结构体中
func ReadConf() (c *conf.Conf) {
	confBytes, err := os.ReadFile(flags.FlagOptions.ConfigFile)
	if err != nil {
		panic(err)
	}
	c = &conf.Conf{}
	err = yaml.Unmarshal(confBytes, c)
	if err != nil {
		panic(fmt.Sprintf("yml配置文件格式错误 %s", err))
	}
	fmt.Printf("读取配置文件%s成功\n", flags.FlagOptions.ConfigFile)
	return
}
