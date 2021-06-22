package tests

import (
	"github.com/streadway/amqp"
)

type MockAmqp struct {
}

func (mock *MockAmqp) Clean() {
}

func (mock *MockAmqp) InitAmqpChannel(amqpServerURL string) {
}

func (mock *MockAmqp) GetAmqpChannel() *amqp.Channel {
	return &amqp.Channel{}
}

func (config *MockAmqp) CloseAmqp() {
}
