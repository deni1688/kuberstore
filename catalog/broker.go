package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

type broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
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

func (b *broker) subscribe() (<-chan amqp.Delivery, error) {
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

func (b *broker) close() error {
	err := new(error)

	*err = b.conn.Close()
	*err = b.ch.Close()

	return *err
}
