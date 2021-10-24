package main

import (
	"log"
)

type listing struct {
	ProductID string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Desc      string `json:"desc,omitempty"`
	ImageUrl  string `json:"imageURL,omitempty"`
}

func main() {
	r, err := newRepo()
	if err != nil {
		log.Fatal("error connecting mongo client", err)
	}
	defer r.close()

	b, err := newBroker()
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
