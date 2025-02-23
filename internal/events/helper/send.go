package helper

import (
	"encoding/json"
	"log"

	rabbitmq "www.github.com/felipegith/internal/rabbit"
)

func Send(queueName string, event interface{}) error {

	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Error to serialize value data", err)
		return err
	}

	err = rabbitmq.Publish(queueName, data)
	if err != nil {
		log.Println("Error to publish event on rabbitmq:", err)
		return err
	}

	log.Println("Event publish successfully ", queueName)
	return nil
}
