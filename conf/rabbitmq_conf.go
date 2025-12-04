package conf

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQConf struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	VirtualHost string `yaml:"virtual-host"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
}

func NewRabbitConn(cfg RabbitMQConf) (*amqp.Connection, error) {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.VirtualHost,
	)
	return amqp.Dial(url)
}
