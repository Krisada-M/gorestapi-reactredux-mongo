package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBclient *mongo.Client = DBconnect()

func errcheck(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DBconnect() *mongo.Client {
	Envload()
	Mongodb := os.Getenv("MONGOURI")
	mongoclient, err := mongo.NewClient(options.Client().ApplyURI(Mongodb))
	errcheck(err)
	ctx, disconnect := context.WithTimeout(context.Background(), 10*time.Second)
	defer disconnect()
	err = mongoclient.Connect(ctx)
	errcheck(err)
	return mongoclient
}

func MongoCollection(client *mongo.Client, collectionname string) *mongo.Collection {
	Envload()
	DBname := os.Getenv("DB_NAME")
	collection := client.Database(DBname).Collection(collectionname)
	return collection
}
