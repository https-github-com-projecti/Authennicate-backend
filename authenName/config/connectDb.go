package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	repo "authenName/repository"
)

func Connect() {
	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	clientOptions := options.Client().ApplyURI(portDB)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	DB := client.Database("authenName")
	repo.CreateCollectionUser(DB)
	repo.CreateCollectionUpload(DB)

	fmt.Println("Connected to MongoDB!")
}
