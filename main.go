package main

// Mongo code to connect to a database and insert queries for benchmarking and performance testing
// using https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandInt(n int) int {
	return 5 + rand.Intn(n)
}

func RandSlice(n int) []string {
	var slice []string
	for i := 0; i < n; i++ {
		slice = append(slice, RandStringRunes(30))
	}
	return slice
}

func InsertPassengers(ctx context.Context, cl *mongo.Collection, n int) {
	for i := 0; i < n; i++ {
		// Insert a single document
		_, err := cl.InsertOne(ctx, bson.D{
			{Key: "name", Value: RandStringRunes(30)},
			{Key: "age", Value: RandInt(100)},
			{Key: "city", Value: RandSlice(RandInt(100))},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("done: ", i)
	}
}

func main() {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Hour)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a collection
	collection := client.Database("test").Collection("train")

	data_set := 800000

	InsertPassengers(ctx, collection, data_set)

	// db.train.find({city:{$in:["uwwy382hw829982wqhg"]}}).explain('executionStats')
	// db.train.find({name:"uwwy382hw829982wqhg"}).explain('executionStats')
	// db.train.count()
	// total city count in train collection: 61271918
	// total document in train collection 1124127

}
