package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// ✅ Correct connection string
	uri := "mongodb+srv://go_user:oYwD3FdakgcegKrV@cluster0.y3fmpu6.mongodb.net/?retryWrites=true&w=majority"

	// Create client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Ping Atlas
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB Atlas:", err)
	}

	fmt.Println("✅ Connected to MongoDB Atlas!")

	// Select DB + Collection
	collection := client.Database("testdb").Collection("users")
	fmt.Println("Using collection:", collection.Name())
}
