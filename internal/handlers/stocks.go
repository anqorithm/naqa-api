package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
	"fmt"
)

type Stock struct {
    Name          string `json:"name" bson:"name"`
    Code          string `json:"code" bson:"code"`
    Sector        string `json:"sector" bson:"sector"`
    ShariaOpinion string `json:"sharia_opinion" bson:"sharia_opinion"`
    Purification  string `json:"purification" bson:"purification"`
}

type StockResponse struct {
    Stocks []Stock `json:"stocks"`
}

func safeString(v interface{}) string {
    if str, ok := v.(string); ok {
        return str
    }
    // Convert float64 to string if it's a number
    if num, ok := v.(float64); ok {
        return fmt.Sprintf("%.2f", num)
    }
    return ""
}

func (h *Handler) GetStocksByYearHandler(c *fiber.Ctx) error {
    collection := h.db.Collection(c.Params("year"))
    
    var stocks []bson.M
    cursor, err := collection.Find(c.Context(), bson.M{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch stocks"})
    }
    defer cursor.Close(c.Context())

    if err := cursor.All(c.Context(), &stocks); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to parse stocks"})
    }

    result := make([]Stock, 0, len(stocks))
    for _, doc := range stocks {
        stock := Stock{
            Name:          safeString(doc["name"]),
            Code:          safeString(doc["code"]),
            Sector:        safeString(doc["sector"]),
            ShariaOpinion: safeString(doc["sharia_opinion"]),
            Purification:  safeString(doc["purification"]),
        }
        result = append(result, stock)
    }
    
    return c.JSON(StockResponse{Stocks: result})
}

func (h *Handler) SearchStocksHandler(c *fiber.Ctx) error {
    filter := bson.M{}
    
    if name := c.Query("name"); name != "" {
        filter["name"] = bson.M{"$regex": name, "$options": "i"}
    }
    if code := c.Query("code"); code != "" {
        filter["code"] = code
    }
    if sector := c.Query("sector"); sector != "" {
        filter["sector"] = sector
    }
    if shariaOpinion := c.Query("sharia_opinion"); shariaOpinion != "" {
        filter["sharia_opinion"] = shariaOpinion
    }
    if purification := c.Query("purification"); purification != "" {
        filter["purification"] = purification
    }
    
    collection := h.db.Collection(c.Params("year"))
    var stocks []bson.M
    cursor, err := collection.Find(c.Context(), filter)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to search stocks"})
    }
    defer cursor.Close(c.Context())

    if err := cursor.All(c.Context(), &stocks); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to parse stocks"})
    }

    result := make([]Stock, 0, len(stocks))
    for _, doc := range stocks {
        stock := Stock{
            Name:          safeString(doc["name"]),
            Code:          safeString(doc["code"]),
            Sector:        safeString(doc["sector"]),
            ShariaOpinion: safeString(doc["sharia_opinion"]),
            Purification:  safeString(doc["purification"]),
        }
        result = append(result, stock)
    }
    
    return c.JSON(StockResponse{Stocks: result})
}