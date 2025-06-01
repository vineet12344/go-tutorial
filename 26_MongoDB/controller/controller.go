package controller

import (
	"context"
	"fmt"
	"log"
	model "mongodbTutorial/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vineetsalve23:Salve_123@gotutorial.vy1hr97.mongodb.net/?retryWrites=true&w=majority&appName=GoTutorial"
const dbName = "netflix"
const colName = "watchlist"

// ! MOST IMPORTANT CONNECTION
var collection *mongo.Collection

// ? Connect with mongoDB

func init() {
	// client optoins
	clientOption := options.Client().ApplyURI(connectionString)

	//connection
	Client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Success")

	collection = Client.Database(dbName).Collection(colName)
	// collection instance is ready

}

// MONGODB HELPERS

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted into DB Sucessfully: ITEM_ID (", inserted.InsertedID, ")")

}

// Update Record

func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

func deleteOneMovie(MovieID string) {
	id, _ := primitive.ObjectIDFromHex(MovieID)
	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count is: l", deleted.DeletedCount)

}

func deleteManyMovies(movieID string) {

	result, err := collection.DeleteMany(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total entries deleted are: ", result.DeletedCount)

}

func findAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie primitive.M

		err := cursor.Decode(movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	fmt.Println(movies)
	return movies
}
