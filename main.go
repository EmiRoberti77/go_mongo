package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Equine represents the data structure for the object to be inserted.
type Equine struct {
    ueln           int    `bson:"ueln"`
    Type           string `bson:"type"`
    Color          string `bson:"color"`
    Vaccination    string `bson:"vaccination"`
    VaccinationDate string `bson:"vaccinationDate"`
}

func main() {
    // MongoDB connection string.
    connectionString := "mongodb+srv://vetapp:vetapp123@cluster0.uwhi5uh.mongodb.net/?retryWrites=true&w=majority"

    // MongoDB database and collection names.
    dbName := "vetapp"
    collectionName := "equines"

    // Create a context with a timeout for database operations.
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Connect to the MongoDB server.
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
    if err != nil {
        fmt.Printf("Error connecting to MongoDB: %v\n", err)
        return
    }

    // Select the database and collection.
    database := client.Database(dbName)
    collection := database.Collection(collectionName)

    // Create an Equine object.
    equine := Equine{
        ueln:           2999909000,
        Type:           "horse type A2",
        Color:          "pink",
        Vaccination:    "completed",
        VaccinationDate: "1/1/2024",
    }

    // Insert the Equine object into the collection.
    _, err = collection.InsertOne(ctx, equine)
    if err != nil {
        fmt.Printf("Error inserting document: %v\n", err)
        return
    }

    fmt.Println("Equine document inserted successfully!")
}
