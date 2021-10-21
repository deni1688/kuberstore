package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	"net/http"
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
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("rabbitmq connection established:", !conn.IsClosed())

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

	r := gin.Default()
	r.Use(corsMiddleware())

	r.POST("/products", func(c *gin.Context) {
		var p product
		_ = json.NewDecoder(c.Request.Body).Decode(&p)

		p.ID = uuid.New().String()
		body, _ := json.Marshal(p)

		err = ch.Publish(
			"products.exchange",
			"products.added",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Println("error publishing message:", err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "product created - id("+p.ID+")",
		})
	})

	log.Fatal(r.Run(":8888"))
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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
