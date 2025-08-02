package services

import (
	"fmt"

	"github.com/Rus203/shop/config"
	"github.com/Rus203/shop/logger"
	"github.com/rabbitmq/amqp091-go"
)

type IMessageConsumer interface {
	DeclareQueue(queueName string) error
	ConsumeEventProcess(queueName string, processor IMessageProcessor) error
}

type MessageConsumer struct {
	conn *configs.RabbitMQConnection
}

func (mc *MessageConsumer) DeclareQueue(queueName string) error {
	return mc.conn.DeclareQueue(queueName)
}

func (mc *MessageConsumer) ConsumeEventProcess(queueName string, processor IMessageProcessor) error {
	channel := mc.conn.GetChannel()

	if channel == nil {
		logger.Panic("messaging channel is nil, retry!")
	}

	msgs, err := channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("error here")
		return fmt.Errorf("failed to consume message: %w", err)
	}

	go func() {
		for msg := range msgs {
			go func(msg amqp091.Delivery) {
				err := processor.ProcessMessage(msg)

				if err != nil {
					logger.Log(fmt.Sprintf("Message processing failed: %v", err))
				}
			}(msg)
		}
	}()

	select {}
}

func NewMessageConsumer() *MessageConsumer {
	rabbitMQConf := configs.GetNewRabbitMQConnection()
	return &MessageConsumer{
		conn: rabbitMQConf,
	}
}
