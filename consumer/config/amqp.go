package config

import (
	"github.com/streadway/amqp"
)

type AmqpConnection interface {
	InitAmqpChannel(amqpServerURL string)
	GetAmqpChannel() *amqp.Channel
	CloseAmqp()
}

type AmqpBroker struct {
	amqpCnx     *amqp.Connection
	amqpChannel *amqp.Channel
}

var localBroker = AmqpBroker{}

func (broker *AmqpBroker) InitAmqpChannel(amqpServerURL string) {
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channelRabbitMQ.QueueDeclare(
		NotifyQueue, // queue name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	localBroker.amqpCnx = connectRabbitMQ
	localBroker.amqpChannel = channelRabbitMQ
}

func (broker *AmqpBroker) GetAmqpChannel() *amqp.Channel {
	return localBroker.amqpChannel
}

func (broker *AmqpBroker) CloseAmqp() {
	localBroker.amqpChannel.Close()
	localBroker.amqpCnx.Close()
}
