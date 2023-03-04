package models

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func Connect() (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {

		log.Fatal(err)
	}
	return &Mongo{
		Client: client,
	}, nil
}

func (db *Mongo) TodoCollection() (todos *mongo.Collection) {
	todos = db.Client.Database("Todo").Collection("Todos")
	return todos
}
