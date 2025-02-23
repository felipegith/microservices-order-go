package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func Init() error {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("Error to connect rabbitmq: %v", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("Error to open a channel: %v", err)
	}

	log.Println("Conected Successfully")
	return nil
}

func Publish(queueName string, message []byte) error {
	if channel == nil {
		return fmt.Errorf("Error channel of rabbitmq not initialize")
	}

	_, err := channel.QueueDeclare(
		queueName, 
		true,      
		false,     
		false,     
		false,     
		nil,       
	)
	if err != nil {
		return fmt.Errorf("Erro to declare a queue '%s': %v", queueName, err)
	}

	err = channel.Publish(
		"",        
		queueName, 
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("Error to publish message on queue '%s': %v", queueName, err)
	}

	log.Printf("Message published '%s' successfully", queueName)
	return nil
}

func Close() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
