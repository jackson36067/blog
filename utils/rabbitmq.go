package utils

import (
	"encoding/json"
	"fmt"

	"blog/consts"
	"blog/models"
	"blog/ws"

	"github.com/streadway/amqp"
)

// SetupRabbitMQ 定义交换机
func SetupRabbitMQ(ch *amqp.Channel) error {
	// 声明延迟交换机
	args := amqp.Table{
		"x-delayed-type": "direct",
	}

	return ch.ExchangeDeclare(
		consts.WsDelayExchangeName, // 交换机名称
		"x-delayed-message",        // 交换机类型
		true,
		false,
		false,
		false,
		args,
	)
}

// SendMessage 用于发送延迟消息
func SendMessage(ch *amqp.Channel, msg models.Message, delayMs int) error {
	body, _ := json.Marshal(msg)

	return ch.Publish(
		consts.WsDelayExchangeName,
		consts.WsRouterKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers: amqp.Table{
				"x-delay": delayMs, // 延迟毫秒
			},
		},
	)
}

// StartConsumer 监听消息
func StartConsumer(ch *amqp.Channel) {
	// 声明队列
	q, _ := ch.QueueDeclare(
		consts.WsQueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	ch.QueueBind(
		q.Name,
		consts.WsRouterKey,
		consts.WsDelayExchangeName,
		false,
		nil,
	)

	messages, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range messages {
			var msg models.Message
			err := json.Unmarshal(d.Body, &msg)

			// websocket 推送
			fmt.Printf("消费消息，推送给用户: %d\n", msg.ReceiveUserID)

			err = ws.WsManager.SendToUser(msg.ReceiveUserID, msg)
			if err != nil {
				fmt.Println("WebSocket 推送失败:", err)
			}
		}
	}()
}
