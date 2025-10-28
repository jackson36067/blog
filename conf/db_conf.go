package conf

import "fmt"

type DBConf struct {
	Mysql MysqlConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
}
type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}
type RedisConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func (m MysqlConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2FShanghai", m.Username, m.Password, m.Host, m.Port, m.Dbname)
}
