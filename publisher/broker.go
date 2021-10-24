package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

type broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func newBroker(user, pass, addr string) (*broker, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s", user, pass, addr)
	conn, err := amqp.Dial(uri)
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

func (b *broker) publish(body []byte) error {
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

func (b *broker) close() error {
	err := new(error)

	*err = b.conn.Close()
	*err = b.ch.Close()

	return *err
}
