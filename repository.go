package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func connect() *mongo.Collection {
	mongoURL := goDotEnvVariable("MONGO_URL")

	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("sw-api").Collection("Planet")

	return collection
}

func getAll() []bson.M {
	collection := connect()

	findOptions := options.Find()
	//findOptions.SetLimit(2)
	findOptions.SetSort(bson.D{{"index", 1}})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func getOneId(id string) bson.M {
	var result bson.M

	collection := connect()

	var val, errrinho = strconv.Atoi(id)
	if errrinho != nil {
		log.Fatal(errrinho)
	}

	var err = collection.FindOne(context.TODO(), bson.D{{"index", val}}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne() ERROR:", err)
		os.Exit(1)
	} else {
		return result
	}
	return bson.M{}
}

func getOneName(name string) bson.M {
	var result bson.M

	collection := connect()

	var err = collection.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne() ERROR:", err)
		os.Exit(1)
	} else {
		return result
	}
	return bson.M{}
}
