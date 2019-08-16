package common

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var DB *mongo.Client

func Init() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), "mongodb+srv://test:test@cluster0-hm9ia.mongodb.net/test?retryWrites=true")
	// client, err := mongo.Connect(context.TODO(), "mongodb://test:test@cluster0-shard-00-00-hm9ia.mongodb.net:27017,cluster0-shard-00-01-hm9ia.mongodb.net:27017,cluster0-shard-00-02-hm9ia.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true")
	// mongodb+srv://test:test@cluster0-hm9ia.mongodb.net/test?retryWrites=true
	// mongodb://test:<password>@cluster0-shard-00-00-hm9ia.mongodb.net:27017,cluster0-shard-00-01-hm9ia.mongodb.net:27017,cluster0-shard-00-02-hm9ia.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	DB = client

	fmt.Println("Connected to MongoDB!")

	return client
}

func Close() {
	err := DB.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	DB = nil

	fmt.Println("Connection to MongoDB closed.")
}

func GetDB() *mongo.Client {
	return DB
}
