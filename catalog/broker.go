package main

import (
	"github.com/streadway/amqp"
	"os"
)

type broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
}

func newBroker() (*broker, error) {
	conn, err := amqp.Dial("amqp://guest:guest@" + os.Getenv("RABBIT_MQ_URI"))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(
		"products.exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		"products.added",
		"products.exchange",
		false,
		nil)
	if err != nil {
		return nil, err
	}

	return &broker{conn, ch, &q}, nil
}

func (b *broker) listen() (<-chan amqp.Delivery, error) {
	return b.ch.Consume(
		b.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}
