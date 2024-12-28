package config

// ###############################################################################
// MongoDB Configuration
// ###############################################################################

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
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGO_DATABASE")
	if dbName == "" {
		dbName = "naqa"
	}

	return &MongoConfig{	
		URI:      uri,
		Database: dbName,
		Timeout:  30 * time.Second, // Increased timeout
	}	
}

// ConnectDB establishes connection to MongoDB
func ConnectDB(cfg *MongoConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(cfg.URI).
		SetTimeout(cfg.Timeout).
		SetConnectTimeout(cfg.Timeout).
		SetSocketTimeout(cfg.Timeout).
		SetRetryWrites(true).
		SetRetryReads(true)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to MongoDB at: %s", cfg.URI)
	return client.Database(cfg.Database), nil
}
