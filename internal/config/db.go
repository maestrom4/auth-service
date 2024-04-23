package config

import (
	"context"

	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var CollectionUser string

func ConnectDB() {
	// Construct the MongoDB URI using environment variables
	username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	database := os.Getenv("MONGO_INITDB_DATABASE")

	CollectionUser = os.Getenv("MONGO_INITDB_COLLECTION")

	// mongoURI := os.Getenv("MONGO_URI")
	log.Info("GraphQL server starting on http://localhost:8082/graphql")
	log.Debugln("username: ", username)
	log.Debugln("password: ", password)
	log.Debugln("database: ", database)
	log.Debugln("collection: ", CollectionUser)
	log.Debugln("mongoURI: ", MongoUrl)

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

	log.Println("Connected to MongoDB!")
}

func GetDBCollection(collectionName string) *mongo.Collection {
	return DB.Database(os.Getenv("MONGO_INITDB_DATABASE")).Collection(collectionName)
}
