package config

import (
	repo "authenName/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Connect() {
	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), 72000*time.Second)
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
		log.Fatal("Can not connect DB url :: " + portDB + " = "  + err.Error())
	}

	DB := client.Database("authnName")
	repo.CreateCollectionUser(DB)
	repo.CreateCollectionUpload(DB)

	fmt.Println("Connected to MongoDB!")
}
