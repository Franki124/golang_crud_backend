package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type HistoryEntry struct {
	Timestamp string `bson:"timestamp"`
	Action    string `bson:"action"`
	Details   string `bson:"details"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       int                `bson:"price"`
	History     []HistoryEntry     `bson:"history"`
}

func NewHistoryEntry(action, details string) HistoryEntry {
	return HistoryEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Action:    action,
		Details:   details,
	}
}

func (p *Product) AddHistoryEntry(action, details string) {
	entry := NewHistoryEntry(action, details)
	p.History = append(p.History, entry)
}

func (p *Product) CreateProduct(client *mongo.Client) error {
	collection := client.Database("Products_db").Collection("Products")
	_, err := collection.InsertOne(context.Background(), p)
	return err
}

func (p *Product) UpdateProduct(client *mongo.Client, action, details string) error {
	p.AddHistoryEntry(action, details)

	collection := client.Database("Products_db").Collection("Products")
	filter := bson.M{"_id": p.ID}
	update := bson.M{
		"$set": bson.M{
			"name":        p.Name,
			"description": p.Description,
			"price":       p.Price,
			"history":     p.History,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func DeleteProduct(client *mongo.Client, productID primitive.ObjectID) error {
	collection := client.Database("Products_db").Collection("Products")

	filter := bson.M{"_id": productID}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}
