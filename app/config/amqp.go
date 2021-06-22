package config

import (
	"github.com/streadway/amqp"
)

type AmqpConnection interface {
	InitAmqpChannel(amqpServerURL string)
	GetAmqpChannel() *amqp.Channel
	Publish(queue, body string) error
	CloseAmqp()
}

type broker struct {
	amqpCnx     *amqp.Connection
	amqpChannel *amqp.Channel
}

var localBroker = broker{}

func (config *ConfigApp) InitAmqpChannel(amqpServerURL string) {
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	localBroker.amqpCnx = connectRabbitMQ
	localBroker.amqpChannel = channelRabbitMQ
}

func (config *ConfigApp) Publish(queue, body string) error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	// Attempt to publish a message to the queue.
	err := localBroker.amqpChannel.Publish(
		"",
		queue,
		false,
		false,
		message,
	)
	if err != nil {
		return err
	}

	return nil
}



func (config *ConfigApp) GetAmqpChannel() *amqp.Channel {
	return localBroker.amqpChannel
}

func (config *ConfigApp) CloseAmqp() {
	localBroker.amqpChannel.Close()
	localBroker.amqpCnx.Close()
}
