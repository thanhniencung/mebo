package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	connStr := fmt.Sprintf("mongodb://host.docker.internal:27017")
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connect database ok")

	Mongo.Client = client
	return Mongo
}