package tests

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/anqorithm/naqa-api/internal/seeders"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
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
func TestLoadDataSourcesConcurrentSeedsOnce(t *testing.T) {
	db := setupTestDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	if err := db.Client().Ping(ctx, nil); err != nil {
		t.Skip("MongoDB not available for testing")
	}
	defer db.Drop(context.Background())

	// The seeder reads the relative "sources" dir, like the API does from the repo root
	wd, _ := os.Getwd()
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(wd)

	var wg sync.WaitGroup
	errs := make(chan error, 3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs <- seeders.LoadDataSources(db)
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("LoadDataSources() error: %v", err)
		}
	}

	count, err := db.Collection("2025").CountDocuments(ctx, bson.M{})
	if err != nil || count != 382 {
		t.Fatalf("2025 count = %d, err = %v, want 382", count, err)
	}
	var aramco bson.M
	if err := db.Collection("2025").FindOne(ctx, bson.M{"code": "2222"}).Decode(&aramco); err != nil {
		t.Fatal(err)
	}
	if aramco["name_en"] != "Saudi Arabian Oil Co." || aramco["logo"] == "" || aramco["sharia_opinion"] != "نقية" {
		t.Fatalf("unexpected aramco document: %v", aramco)
	}
}
