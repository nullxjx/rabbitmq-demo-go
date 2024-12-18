package main

import (
	"fmt"
	"log"
	"rabbitmq-demo-go/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
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

	//err = ch.Qos(
	//	1,     // 预取值
	//	0,     // 预取大小（0 表示不限制大小）
	//	false, // 是否全局（false 表示只对当前消费者生效）
	//)

	// 消费消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者标签
		false,  // 自动确认
		false,  // 排他性
		false,  // 无本地
		false,  // 无等待
		nil,    // 参数
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// 每隔 2 秒消费一条消息
	for msg := range msgs {
		fmt.Printf("Received: %s\n", msg.Body)
		//time.Sleep(1 * time.Second) // 每隔 2 秒消费一条消息
		time.Sleep(200 * time.Millisecond) // 每隔 2 秒消费一条消息
		if err := msg.Ack(false); err != nil {
			log.Fatalf("Failed to ack a message: %v", err)
		}
	}
}
