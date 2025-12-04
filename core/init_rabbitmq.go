package core

import (
	"blog/conf"
	"blog/global"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() *amqp.Connection {
	rabbitMQ := global.Conf.RabbitMQ
	conn, err := conf.NewRabbitConn(rabbitMQ)
	if err != nil {
		logrus.Fatalf("RabbitMQ 连接失败: %v", err)
	}
	logrus.Infof("rabbitmq连接成功")
	return conn
}
