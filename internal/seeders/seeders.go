package seeders

// ###############################################################################
// Stock Data Seeder
// ###############################################################################

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	sourcesDir = "sources"
)

type StockData struct {
	ID            int    `json:"id" bson:"id"`
	Purification  string `json:"purification" bson:"purification"`
	ShariaOpinion string `json:"sharia_opinion" bson:"sharia_opinion"`
	Name          string `json:"name" bson:"name"`
	Code          string `json:"code" bson:"code"`
	Sector        string `json:"sector" bson:"sector"`
	DataSource    string `json:"data_source" bson:"data_source"`
}

type JSONResponse struct {
	Stocks []StockData `json:"stocks"`
}

func getCollectionName(filename string) string {
	return strings.TrimSuffix(strings.TrimPrefix(filepath.Base(filename), "final_"), ".json")
}

func LoadDataSources(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	files, err := filepath.Glob(filepath.Join(sourcesDir, "final_*.json"))
	if err != nil {
		return fmt.Errorf("error finding JSON files: %v", err)
	}

	for _, file := range files {
		collectionName := getCollectionName(file)
		collection := db.Collection(collectionName)

		count, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			return fmt.Errorf("error checking collection %s: %v", collectionName, err)
		}

		if count > 0 {
			fmt.Printf("Data already exists in collection %s, skipping\n", collectionName)
			continue
		}

		fmt.Printf("Processing file: %s into collection: %s\n", file, collectionName)

		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", file, err)
		}

		var jsonResponse JSONResponse
		if err := json.Unmarshal(data, &jsonResponse); err != nil {
			return fmt.Errorf("error parsing JSON from %s: %v", file, err)
		}

		var documents []interface{}
		for _, stock := range jsonResponse.Stocks {
			stock.DataSource = filepath.Base(file)
			documents = append(documents, stock)
		}

		if len(documents) > 0 {
			_, err = collection.InsertMany(ctx, documents)
			if err != nil {
				return fmt.Errorf("error inserting documents into %s: %v", collectionName, err)
			}
			fmt.Printf("Successfully inserted %d documents into collection %s\n", len(documents), collectionName)
		}
	}

	fmt.Println("Data seeding completed successfully")
	return nil
}
