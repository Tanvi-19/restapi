package configuration

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Collection {
	ctx, exit := context.WithTimeout(context.TODO(), 10*time.Second)
	defer exit()

	client := options.Client().ApplyURI("mongodb://192.168.0.102:27017")

	c, err := mongo.Connect(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} else {
		fmt.Println("Connected to mongodb")
	}

	return c.Database("Employees").Collection("emp")
}