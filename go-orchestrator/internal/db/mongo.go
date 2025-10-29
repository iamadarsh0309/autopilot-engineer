package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(" Mongo connection failed:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(" Mongo ping failed:", err)
	}

	MongoClient = client
	log.Println(" Connected to MongoDB successfully")
}

func GetCollection(dbName, colName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(colName)
}
