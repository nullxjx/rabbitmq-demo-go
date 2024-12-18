package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq-demo-go/config"
)

func main() {
	// 连接到 RabbitMQ 服务器
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.User, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		config.Queue, // 队列名称
		false,        // 持久化
		false,        // 自动删除
		false,        // 排他性
		false,        // 无等待
		nil,          // 参数
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 生产者每隔 1 秒发送一条消息
	for {
		timestamp := time.Now().Format(time.RFC3339) // 获取当前时间戳
		body := fmt.Sprintf("Message sent at: %s", timestamp)

		// 发送消息
		err = ch.Publish(
			"",     // 交换机
			q.Name, // 路由键
			false,  // 强制性
			false,  // 立即
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			log.Printf("Failed to publish a message: %v", err)
		} else {
			fmt.Printf("Sent: %s\n", body)
		}

		time.Sleep(1 * time.Second) // 每隔 1 秒发送一条消息
	}
}
