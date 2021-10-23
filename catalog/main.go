package main

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type listing struct {
	ProductID string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Desc      string `json:"desc,omitempty"`
	ImageUrl  string `json:"imageURL,omitempty"`
}

func main() {
	msgs, mqConn, err := getMQConn()
	if err != nil {
		log.Fatal("error consuming queue messages:", err)
	}

	defer mqConn.Close()

	client, cancel, err := getMongoConn()
	if err != nil {
		log.Fatalln("error connecting mongo client", err)
	}
	defer cancel()

	item := new(listing)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			_ = json.Unmarshal(d.Body, &item)
			log.Printf("New products event of type %s with id %s", d.RoutingKey, item.ProductID)
			createListing(client, item)
		}
	}()

	log.Printf(" [*] Waiting for product events. To exit press CTRL+C")
	<-forever
}

func createListing(client *mongo.Client, item *listing) {
	collection := client.Database("catalog").Collection("listings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	res, err := collection.InsertOne(ctx, bson.M{
		"productId": item.ProductID,
		"name":      item.Name,
		"desc":      item.Desc,
		"imageUrl":  item.ImageUrl,
	})
	if err != nil {
		log.Println("failed to insert item", err)
	} else {
		log.Println("created new catalog listing with", res.InsertedID.(primitive.ObjectID).Hex())
	}
}

func getMongoConn() (*mongo.Client, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin123@"+os.Getenv("MONGO_URI")+"/catalog"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalln("error disconnecting mongo client", err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("error pinging database", err)
	}

	log.Println("mongo connection successful!")

	return client, cancel, err
}

func getMQConn() (<-chan amqp.Delivery, *amqp.Connection, error) {
	conn, err := connect()
	if err != nil {
		log.Fatal("error getting connection", err)
	}

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
	return msgs, conn, err
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
