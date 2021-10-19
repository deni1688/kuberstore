package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type product struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Desc     string `json:"desc,omitempty"`
	ImageUrl string `json:"imageURL,omitempty"`
	Stock    int    `json:"stock,omitempty"`
}

func main() {
	conn, err := connect()
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("error getting connection", err)
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(
		"products.exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatal("error initializing exchange", err)
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
		log.Fatal("error initializing a queue")
	}

	err = ch.QueueBind(
		q.Name,
		"products.added",
		"products.exchange",
		false,
		nil)
	if err != nil {
		log.Fatal("error binding a queue", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("error consuming queue messages:", err)
	}

	p := new(product)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			_ = json.Unmarshal(d.Body, &p)
			log.Printf("New products event of type %s with id %s", d.RoutingKey, p.ID)
		}
	}()

	log.Printf(" [*] Waiting for product events. To exit press CTRL+C")
	<-forever
}

func connect() (*amqp.Connection, error) {
	retries := 6

	var conn *amqp.Connection
	var err error

	for i := 0; i < retries; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@" + os.Getenv("RABBIT_MQ_URI"))
		if err == nil {
			return conn, err
		}
		time.Sleep(1 * time.Second)
	}

	return conn, err
}

