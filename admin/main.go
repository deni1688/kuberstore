package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"io"
	"log"
	"net/http"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		log.Fatalln(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)

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
		log.Fatal("error initializing exchange", err)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/products", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)

		_ = ch.Publish(
			"products.exchange",
			"products.added"	,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)

		c.JSON(http.StatusOK, gin.H{
			"message": "event published",
		})
	})

	log.Fatal(r.Run(":8888"))
}

func CORSMiddleware() gin.HandlerFunc {
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
