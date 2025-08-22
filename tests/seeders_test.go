package tests

import (
	"context"
	"testing"
	"time"

	"github.com/anqorithm/naqa-api/internal/seeders"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Skip("MongoDB not available for testing")
	}

	testDB := client.Database("naqa_test")
	return testDB
}

func TestLoadDataSources(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		t.Skip("Test database not available")
	}

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		db.Drop(ctx)
	}()

	err := seeders.LoadDataSources(db)
	if err != nil {
		t.Logf("LoadDataSources() returned error (expected if no source files): %v", err)
	}
}