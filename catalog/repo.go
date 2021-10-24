package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type repo struct {
	client *mongo.Client
}

func newRepo() (*repo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://admin:admin123@" + os.Getenv("MONGO_URI") + "/catalog?authSource=admin"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Println("mongo connection successful!")

	return &repo{client}, nil
}



func (d *repo) insert(item listing) (string, error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	collection := d.client.Database("catalog").Collection("listings")
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.M{
		"productId": item.ProductID,
		"name":      item.Name,
		"desc":      item.Desc,
		"imageUrl":  item.ImageUrl,
	})
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

