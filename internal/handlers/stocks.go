package handlers

import (
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func cleanNaNValues(doc bson.M) bson.M {
	for k, v := range doc {
		switch value := v.(type) {
		case float64:
			if math.IsNaN(value) {
				doc[k] = nil
			}
		}	
	}
	return doc
}

func (h *Handler) GetStocksByYearHandler(c *fiber.Ctx) error {
	year := c.Params("year")
	
	collection := h.db.Collection(year)
	fmt.Println(collection)
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch stocks",
		})
	}
	
	var stocks []bson.M
	if err = cursor.All(c.Context(), &stocks); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to parse stocks",
		})
	}

	// Clean NaN values
	for i := range stocks {
		stocks[i] = cleanNaNValues(stocks[i])
	}
	
	return c.JSON(fiber.Map{
		"stocks": stocks,
	})
}

func (h *Handler) SearchStocksHandler(c *fiber.Ctx) error {
	// Get year from URL parameters
	year := c.Params("year")
	
	// Get query parameters
	name := c.Query("name")
	code := c.Query("code")
	sector := c.Query("sector")
	shariaOpinion := c.Query("sharia_opinion")
	
	// Build filter
	filter := bson.M{}
	
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	if code != "" {
		filter["code"] = code
	}
	if sector != "" {
		filter["sector"] = sector
	}
	if shariaOpinion != "" {
		filter["sharia_opinion"] = shariaOpinion
	}
	
	collection := h.db.Collection(year)
	cursor, err := collection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to search stocks",
		})	
	}
	
	var stocks []bson.M
	if err = cursor.All(c.Context(), &stocks); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to parse stocks",
		})
	}

	// Clean NaN values
	for i := range stocks {
		stocks[i] = cleanNaNValues(stocks[i])
	}
	
	return c.JSON(fiber.Map{
		"stocks": stocks,
	})
}
