package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	mongoURI = "mongodb://mongo:27017"
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "50001"
)

type Config struct {
}

var client *mongo.Client

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

}
func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "root",
	})

	// connect to mongo
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting:", err)
		return nil, err
	}

	return c, nil
}
