package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017" // To use environment variables for better security
const dbName = "netflix"                             // Database name
const collectionName = "movies"                      // Collection name required for mongo kinda like tables in SQL

// Important collection object
var collection *mongo.Collection

// Connects to MongoDB
// init is a special function that runs before main only once when app starts
func init() {
	// Set client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	// context is a package in golang. It is used when we make calls to external resources like databases
	// TODO is used when we don't know what to pass in context
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Set collection
	// Syntax: client.Database("dbName").Collection("collectionName")
	collection = client.Database(dbName).Collection(collectionName)

	// if collection ready, print message
	fmt.Println("Collection instance ready!")
}
