package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client
var dbCollection *mongo.Collection
var ctx context.Context

// initDB : instantiate client and collection global variables
func initDB() {

	// ctx : replaces context.TODO() because it's confusing
	ctx := context.Background()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	dbClient, err := mongo.Connect(ctx, clientOptions)
	checkPanic(err)

	// Check the connection
	err = dbClient.Ping(ctx, nil)
	checkPanic(err)

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	dbCollection = dbClient.Database("test").Collection("products")
}

// disconnectDB : lose the connection once no longer needed
func disconnectDB() {
	err := dbClient.Ping(ctx, nil)
	checkPanic(err)
	fmt.Println("Successfully pinged before disconnection.")

	err = dbClient.Disconnect(ctx)
	checkPanic(err)
	fmt.Println("Connection to MongoDB closed.")
}

// eraseDB : delete all entries in the DB. Useful at the end of tests
func eraseDB(verbose bool) {
	deleteResult, err := dbCollection.DeleteMany(ctx, bson.M{})
	checkPanic(err)
	if verbose {
		fmt.Printf("Deleted %v documents in the products collection\n", deleteResult.DeletedCount)
	}
}
