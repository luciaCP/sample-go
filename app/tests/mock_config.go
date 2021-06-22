package tests

import (
	"github.com/streadway/amqp"
)

type MockAmqp struct {
	PublishTimesCalled int
}

func (mock *MockAmqp) Clean() {
	mock.PublishTimesCalled = 0
}

func (mock *MockAmqp) InitAmqpChannel(amqpServerURL string) {
}

func (mock *MockAmqp) Publish(queue, body string) error {
	mock.PublishTimesCalled++
	return nil
}

func (mock *MockAmqp) GetAmqpChannel() *amqp.Channel {
	return &amqp.Channel{}
}

func (config *MockAmqp) CloseAmqp() {
}