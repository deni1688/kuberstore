package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type product struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Desc     string `json:"desc,omitempty"`
	ImageUrl string `json:"imageURL,omitempty"`
	Stock    int    `json:"stock,omitempty"`
}

func main() {
	b, err := newBroker()
	if err != nil {
		log.Fatal("error connecting to broker", err)
	}
	defer b.close()

	r := gin.Default()
	r.Use(cors())

	r.POST("/products", getCreateProductHandler(b))

	log.Fatal(r.Run(":8888"))
}
