package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB() (*mongo.Client, error) {
	// Set connection options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func executeMongoQuery(client *mongo.Client, collectionName string, query interface{}) error {
	// Access the specified collection
	collection := client.Database("your_database_name").Collection(collectionName)

	// Execute the query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, query)
	if err != nil {
		return err
	}

	fmt.Println("Query executed successfully!")
	return nil
}

func mongoSample() {
	// Connect to MongoDB
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	// Example query execution
	query := bson.M{"name": "John Doe"}
	err = executeMongoQuery(client, "users", query)
	if err != nil {
		log.Fatal(err)
	}
}
