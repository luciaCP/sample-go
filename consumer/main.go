package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sample-go/consumer/callbacks"
	"sample-go/consumer/config"
)

func main() {
	config.Connections.InitDb("postgresql://postgres@0.0.0.0:5432", "go_test")

	connectRabbitMQ, err := amqp.Dial("amqp://localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	messages, err := channelRabbitMQ.Consume(
		"IncreaseQueue", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			callbacks.Interpret(fmt.Sprintf("%s", message.Body))
		}
	}()

	<-forever
}
