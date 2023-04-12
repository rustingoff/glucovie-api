package mongodb

import (
	"context"
	"glucovie/pkg/dotenv"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBConnection() *mongo.Database {
	connectionURI := dotenv.GetEnvironmentVariable("MONGO_CONN_URI")
	if connectionURI == "" {
		log.Fatalf("MONGO_CONN_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(connectionURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("mogno connection failure: %v", err)
	}

	dbName := dotenv.GetEnvironmentVariable("MONGODB_NAME")
	if dbName == "" {
		log.Fatalf("MONGODB_NAME is not set")
	}
	db := client.Database(dbName)

	return db
}
