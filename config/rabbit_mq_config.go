package configs

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQConnection struct {
	conn *amqp091.Connection
	queue string
}


func (rc *RabbitMQConnection) Reconnect() *amqp091.Connection {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d", Env.RabbitMQUsername, Env.RabbitMQPassword, Env.RabbitMQHost, Env.RabbitMQPort)
	conn, err := amqp091.Dial(amqpURL)

	if err != nil {
		panic(err)
	}

	log.Println("rabbitmq has been reconnected")

	return conn
}

func (rc *RabbitMQConnection) DeclareQueue(queueName string) error {
	channel, err := rc.conn.Channel()

	if err != nil {
		panic(err)
	}

	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}

func (rc *RabbitMQConnection) GetConnection() *amqp091.Connection {
	if rc.conn == nil {
		return rc.Reconnect()
	}

	return rc.conn
}

func (rc *RabbitMQConnection) GetChannel() *amqp091.Channel {
	if rc.conn == nil {
		rc.Reconnect()
	}

	channel, err := rc.conn.Channel()

	if err != nil {
		panic(err)
	}

	if channel != nil && channel.IsClosed() {
		log.Println("channel was closed and error creating channel")
	}

	return channel
}

func (rc *RabbitMQConnection) GetQueue() string {
	return rc.queue
}


func GetNewRabbitMQConnection() *RabbitMQConnection {
	defaultQueueName := Env.RabbitMQDefaultQueue

	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d", Env.RabbitMQUsername, Env.RabbitMQPassword, Env.RabbitMQHost, Env.RabbitMQPort)
	conn, err := amqp091.Dial(amqpURL)

	if err != nil {
		panic(err)
	}


	return &RabbitMQConnection{
		queue: defaultQueueName,
		conn: conn,
	}
}