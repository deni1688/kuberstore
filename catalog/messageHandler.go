package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func messageHandler(r *repo, msg amqp.Delivery) {
	var item listing

	err := json.Unmarshal(msg.Body, &item)
	if err != nil {
		log.Println("error parsing message body", err)
		return
	}

	log.Printf("new products event of type %s with id %s", msg.RoutingKey, item.ProductID)

	id, err := r.insert(item)
	if err != nil {
		log.Println("error inserting item", err)
	} else {
		log.Println("inserted catalog listing with id: " + id)
	}
}
