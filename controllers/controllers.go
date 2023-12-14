package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/generalpy101/Go-Mongo-API/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://mongodb:27017" // To use environment variables for better security, using dockerized mongodb
const dbName = "netflix"                           // Database name
const collectionName = "movies"                    // Collection name required for mongo kinda like tables in SQL

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

// MongoDB helper - TODO: Move to a separate file

// Insert one movie into DB
// Helper so we don't have to export this function
func insertOneMovie(movie models.Netflix) {
	res, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Inserted one movie with id: ", res.InsertedID) // InsertedID is a special field in mongo
}

// Update one movie in DB
func updateOneMovie(movieID string) {
	mongoId, err := primitive.ObjectIDFromHex(movieID) // Convert string to mongo ObjectID

	if err != nil {
		log.Fatal(err)
		return
	}
	// Everything is bson in mongo so filter has to be BSON too
	// Read docs when lost
	// bson.M is used when we want shorter and clearer filter declaration and order doesn't matter
	// If we have compound filter, we use bson.D
	filter := bson.M{"_id": mongoId} // Filter to find the movie

	// Update fields
	// creating filter to update
	// Some mongodb syntax
	update := bson.M{
		"$set": bson.M{
			// Setting moview to watched true for now, for learning purposes
			"watched": true,
		},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
}

// Delete one movie in DB
func deleteOneMovie(movieID string) {
	mongoId, err := primitive.ObjectIDFromHex(movieID) // Convert string to mongo ObjectID

	if err != nil {
		log.Fatal(err)
		return
	}
	// Everything is bson in mongo so filter has to be BSON too
	// Read docs when lost
	// bson.M is used when we want shorter and clearer filter declaration and order doesn't matter
	// If we have compound filter, we use bson.D
	filter := bson.M{"_id": mongoId} // Filter to find the movie

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Deleted documents count: ", result.DeletedCount)
}

// Delete all movies
func deleteAllMovies() int64 {
	// Delete all documents in collection
	// Empty filter means delete everything
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	return result.DeletedCount
}

// Get all movies
func getAllMovies() []primitive.M {
	// Here cursor is returned instead of result
	// Cursor is like a pointer to the result, will have to iterate over it, close it too
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M // slice of bson.M

	// Close cursor when done
	defer cursor.Close(context.Background())

	// Iterate over cursor
	// cursor.Next() returns true if there is a next document, context is required
	for cursor.Next(context.Background()) {
		var movie bson.M // bson.M is a map, a key value pair

		// TIP: can define variables and use them in if statement
		/*
			eg:
			if err := cursor.Decode(&movie); err != nil {
				log.Fatal(err)
			}
		*/

		err := cursor.Decode(&movie) // Decode current document to movie
		if err != nil {
			// Fatal is followed by os.Exit(1) so no need to return, change old code
			log.Fatal(err)
		}

		movies = append(movies, movie) // Append movie to movies slice
	}

	return movies
}

// Controllers to be used in routers - This would be kept in this file
func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies := getAllMovies()

	// Write movies to response
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		// TODO: Better error handling for API
		log.Fatal(err)
	}
}

// Create a movie
func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// Create movie object
	var movie models.Netflix

	// Decode json to movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	// Insert movie to DB
	insertOneMovie(movie)

	// Write movie to response
	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		log.Fatal(err)
	}
}

// Set movie as watched
func MarkMovieAsWatchedController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	pramas := mux.Vars(r) // Get params
	movieID := pramas["id"]

	updateOneMovie(movieID)

	// Write movies to response
	err := json.NewEncoder(w).Encode(`{
		"message": "Movie marked as watched"
	}`)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete a movie
func DeleteMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	pramas := mux.Vars(r) // Get params
	movieID := pramas["id"]

	deleteOneMovie(movieID)

	// Write movies to response
	err := json.NewEncoder(w).Encode(`{
		"message": "Moview with id ` + movieID + ` deleted"
	}`)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete all movies
func DeleteAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deletedCount := deleteAllMovies()

	// Write movies to response
	err := json.NewEncoder(w).Encode(`{
		"message": "Deleted ` + fmt.Sprint(deletedCount) + ` movies"
	}`)
	if err != nil {
		log.Fatal(err)
	}
}
