package core

import (
	"blog/conf"
	"blog/global"
	"blog/utils"

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
	ch, err := conn.Channel()
	if err != nil {
		logrus.Fatalf("获取 channel 失败：" + err.Error())
	}
	// 声明交换机和队列
	if err := utils.SetupRabbitMQ(ch); err != nil {
		logrus.Fatalf("创建交换机失败" + err.Error())
	}
	// 启动消费者
	utils.StartConsumer(ch)
	return conn
}
