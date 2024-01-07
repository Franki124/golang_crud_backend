package controller

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("No MongoDB URI provided in environment variables")
		return nil, errors.New("no MongoDB URI provided in environment variables")
	}

	opts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal("Failed to create client:", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")
	return client, nil
}

func DisconnectFromDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}

	log.Println("Successfully disconnected from MongoDB")
}
