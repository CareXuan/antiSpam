package base

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	key string
	//连接信息
	Mqurl string
}

func InitChannel(queue string) *amqp.Channel {
	rabbitmq := &RabbitMQ{QueueName: queue, Exchange: "carexuan_exchange", key: "testkey", Mqurl: "amqp://guest:guest@localhost:5672/"}
	//创建RabbitMQ连接
	rabbitmq.conn, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
	channel, _ := rabbitmq.conn.Channel()
	_, err := channel.QueueDeclare(
		"carexuan_test",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil
	}
	return channel
}

func Publish(channel *amqp.Channel,content string) error{
	err := channel.Publish(
		"",
		"carexuan_test",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(content),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
