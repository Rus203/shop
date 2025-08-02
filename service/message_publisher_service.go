package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rabbitmq/amqp091-go"

	"github.com/Rus203/shop/config"

	"github.com/Rus203/shop/logger"
)

type IMessagePublisher interface {
	PublishEvent(queueName string, body any) error
	DeclareQueue(queueName string) error
}

type MessagePublisher struct {
	conn *configs.RabbitMQConnection
}

func (mp *MessagePublisher) PublishEvent(queueName string, body any) error {
	json, err := json.Marshal(body)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if queueName == "" {
		queueName = configs.Env.RabbitMQDefaultQueue
	}

	channel := mp.conn.GetChannel()

	if channel == nil {
		logger.Panic("messaging channel is nil, retry !")
	}

	if channel.IsClosed() {
		logger.Panic("could not publish event, channel closed")
	}

	return channel.PublishWithContext(
		ctx, "", queueName, false, false, amqp091.Publishing{
			ContentType:  "application/json",
			Body:         json,
			DeliveryMode: amqp091.Persistent,
		},
	)
}

func (mp *MessagePublisher) DeclareQueue(queueName string) error {
	return mp.conn.DeclareQueue(queueName)
}

func NewMessagePublisher() *MessagePublisher {
	rabbitMQConf := configs.GetNewRabbitMQConnection()

	rabbitMQConf.DeclareQueue(rabbitMQConf.GetQueue())

	return &MessagePublisher{
		conn: rabbitMQConf,
	}
}
