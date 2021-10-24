package main

import (
	"log"
	"os"
)

type listing struct {
	ProductID string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Desc      string `json:"desc,omitempty"`
	ImageUrl  string `json:"imageURL,omitempty"`
}

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbAddr = os.Getenv("DB_ADDR")
	mqUser = os.Getenv("MQ_USER")
	mqPass = os.Getenv("MQ_PASS")
	mqAddr = os.Getenv("MQ_ADDR")
)

func main() {
	r, err := newRepo(dbUser, dbPass, dbAddr, "catalog")
	if err != nil {
		log.Fatal("error connecting mongo client", err)
	}
	defer r.close()

	b, err := newBroker(mqUser, mqPass, mqAddr)
	if err != nil {
		log.Fatal("error connecting broker", err)
	}
	defer b.close()

	msgs, err := b.subscribe()
	if err != nil {
		log.Fatal("error getting messages", err)
	}

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			messageHandler(r, m)
		}
	}()

	log.Printf(" [*] Waiting for events")
	<-forever
}
