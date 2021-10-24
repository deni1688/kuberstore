package main

import (
	"github.com/streadway/amqp"
	"os"
)

type broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
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

	return &broker{conn, ch}, nil
}

func (b *broker) publish(body []byte) error  {
	return b.ch.Publish(
		"products.exchange",
		"products.added",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
