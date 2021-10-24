package main

import (
	"encoding/json"
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


	b, err := newBroker()
	if err != nil {
		log.Fatal("error connecting broker", err)
	}

	msgs, err := b.listen()
	if err != nil {
		log.Fatal("error getting messages", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var item listing
			_ = json.Unmarshal(d.Body, &item)

			log.Printf("new products event of type %s with id %s", d.RoutingKey, item.ProductID)

			id, err := r.insert(item)
			if err != nil {
				log.Println("error inserting item", err)
			} else {
				log.Println("inserted catalog listing with id: " + id)
			}
		}
	}()

	log.Printf(" [*] Waiting for events")
	<-forever
}
