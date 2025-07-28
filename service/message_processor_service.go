package services

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"


	"github.com/Rus203/shop/constants"
	"github.com/Rus203/shop/logger"
	"github.com/Rus203/shop/util"

)

type IMessageProcessor interface{
	ProcessMessage(message any) error
}

type MessageProcessor struct {
	publisher IMessagePublisher;
	connection *map[string]IWebSocketConnection
	mutex sync.RWMutex
}

func (mp *MessageProcessor) ProcessMessage(message any) error {
	msg := message.(amqp091.Delivery)
	var event map[string]any

	if err := json.Unmarshal(msg.Body, &event); err != nil {
		msg.Nack(false, true)
		return err
	}

	logger.Log(fmt.Sprintf("processing message: %v", event))

	if value, ok := event["order_status"]; ok {
		var err error
		switch value {
			case constants.ORDER_ORDERED: {
				err = mp.handleOrderOrdered(event)
			}

			case constants.ORDER_PREPARING: {
				err = mp.handleOrderPreparing(event)
			}

			case constants.ORDER_PREPARED: {
				err = mp.handleOrderPrepared(event)
			}

			default: {
				logger.Log("No order to be processed!")
			}
		}

		if err != nil {
			msg.Nack(false, true)
			logger.Log(fmt.Sprintf("Error Processing Message: %v", err))
			return err
		}
	}

	return msg.Ack(false)
}

func (mp *MessageProcessor) handleOrderOrdered(event map[string]any) error {
	logger.Log(fmt.Sprintf("order %v accepted", event))
	event["order_status"] = constants.ORDER_PREPARING

	if err := mp.publisher.PublishEvent(constants.KITCHEN_ORDER_QUEUE, event); err != nil {
		logger.Log(fmt.Sprintf("error: %v, event: %v", err, event))

		message := map[string]string {
			"order_status": "",
			"error":   err.Error(),
		}

		if mp.connection != nil {
			json, err := json.Marshal(message)

			if err != nil {
				return err
			}

			mp.mutex.Lock()
			defer mp.mutex.Unlock()

			pizza, ok := (*mp.connection)["pizza"]	 // todo: implement multiply conevtion storing

			if  ok && pizza != nil {
				err = (*mp.connection)["pizza"].SendMessage(json)
			}

			return err
		}

		return err
	}

	return nil
}

func (mp *MessageProcessor) handleOrderPreparing(event map[string]any) error {
	logger.Log(fmt.Sprintf("order %v accepted", event))
	event["order_status"] = constants.ORDER_PREPARED

	time.Sleep(utils.GenerateRandomDuration(1, 6))	// todo: add bl here later

	if err := mp.publisher.PublishEvent(constants.KITCHEN_ORDER_QUEUE, event); err != nil {
		logger.Log(fmt.Sprintf("error: %v, event: %v", err, event))

		message := map[string]string {
			"order_status": "",
			"error":   err.Error(),
		}

		if mp.connection != nil {
			json, err := json.Marshal(message)

			if err != nil {
				return err
			}
		
			mp.mutex.Lock()
			defer mp.mutex.Unlock()

			pizza, ok := (*mp.connection)["pizza"]	 // todo: implement multiply conevtion storing

			if  ok && pizza != nil {
				err = (*mp.connection)["pizza"].SendMessage(json)
			}

			return err
		}

		return err
	}

	return nil
}

func (mp *MessageProcessor) handleOrderPrepared(event map[string]any) error {
	logger.Log(fmt.Sprintf("order %v accepted", event))
	event["order_status"] = constants.ORDER_DELIVERED

	time.Sleep(utils.GenerateRandomDuration(1, 6))	// todo: add bl here later

	logger.Log(fmt.Sprintf("order %v prepared successfully", event["order_no"]))

	message := map[string]any{
		"message": constants.ORDER_PREPARED_SUCCESSFULLY,
		"order":   event,
	}

	if mp.connection != nil {
		json, err := json.Marshal(message)

		if err != nil {
			return err
		}

		mp.mutex.Lock()
		defer mp.mutex.Unlock()

		pizza, ok := (*mp.connection)["pizza"]	 // todo: implement multiply conevtion storing

		if  ok && pizza != nil {
			err = (*mp.connection)["pizza"].SendMessage(json)
		}

		return err
	}

	return nil
}

func NewMessageProcessor(publisher IMessagePublisher, connection *map[string]IWebSocketConnection) *MessageProcessor {
	return &MessageProcessor{
		publisher: publisher,
		connection: connection,
	}
}
