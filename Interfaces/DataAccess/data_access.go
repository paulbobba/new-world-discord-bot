package dataaccess

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoDBConnectionStringEnvVarName = "MONGODB_CONNECTION_STRING"
const MongoDBDatabaseStringEnvVarName = "MONGODB_DATABASE"
const MongoDBCollectionStringEnvVarName = "MONGODB_COLLECTION"

// connects to MongoDB
func NewClient(mongoDBConnectionString string) *mongo.Client {

	// mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		log.Fatal("missing environment variable: ", mongoDBConnectionString)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBConnectionString).SetDirect(true)
	c, err := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	return c
}
