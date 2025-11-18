package conf

type Conf struct {
	System SystemConf `yaml:"system"`
	DB     DBConf     `yaml:"db"`
	Email  EmailConf  `yaml:"email"`
	Log    LogConf    `yaml:"log"`
	AliOss AliOssConf `yaml:"ali_oss"`
}
