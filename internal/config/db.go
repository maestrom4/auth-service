package config

import (
	"context"
	"fmt"

	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {

	serverURL := fmt.Sprintf("http://localhost:%d/graphql", MongoUsersPort)
	log.Debugln("GraphQL server starting on", serverURL)

	// Connect to MongoDB
	var err error
	DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoUrl))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = DB.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Infoln("Connected to MongoDB!")
}

func GetDBCollection(collectionName string) *mongo.Collection {
	return DB.Database(Database).Collection(collectionName)
}
