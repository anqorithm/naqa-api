package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig holds MongoDB configuration
type MongoConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

// NewMongoConfig creates a new MongoDB configuration
func NewMongoConfig() *MongoConfig {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		dbName = "naqa"
	}

	return &MongoConfig{
		URI:      uri,
		Database: dbName,
		Timeout:  10 * time.Second,
	}
}

// ConnectDB establishes connection to MongoDB
func ConnectDB(cfg *MongoConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client.Database(cfg.Database), nil
}
