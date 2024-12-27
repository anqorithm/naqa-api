package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

type PurificationRequest struct {
    StartDate    string  `json:"start_date"`
    EndDate      string  `json:"end_date"`
    NumberOfStocks int   `json:"number_of_stocks"`
    StockCode    string  `json:"stock_code"`
}

type PurificationResponse struct {
    PurificationAmount float64 `json:"purification_amount"`
    DaysHeld          int     `json:"days_held"`
    PurificationRate  float64 `json:"purification_rate"`
}

func safeString(v interface{}) string {
    if str, ok := v.(string); ok {
        return str
    }
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
        return SendError(c, fiber.StatusInternalServerError, ErrCodeDatabaseError, 
            "Failed to fetch stocks", nil)
    }
    defer cursor.Close(c.Context())

    if err := cursor.All(c.Context(), &stocks); err != nil {
        return SendError(c, fiber.StatusInternalServerError, ErrCodeDatabaseError, 
            "Failed to parse stocks", nil)
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

func (h *Handler) CalculatePurificationHandler(c *fiber.Ctx) error {
    var req PurificationRequest
    if err := c.BodyParser(&req); err != nil {
        return SendError(c, fiber.StatusBadRequest, ErrCodeInvalidRequest, 
            "Invalid request body", nil)
    }

    startDate, err := time.Parse("2006-01-02", req.StartDate)
    if err != nil {
        return SendError(c, fiber.StatusBadRequest, ErrCodeInvalidDateFormat, 
            "Invalid start date format", map[string]string{
                "expected_format": "YYYY-MM-DD",
                "provided_value": req.StartDate,
            })
    }

    endDate, err := time.Parse("2006-01-02", req.EndDate)
    if err != nil {
        return SendError(c, fiber.StatusBadRequest, ErrCodeInvalidDateFormat, 
            "Invalid end date format", map[string]string{
                "expected_format": "YYYY-MM-DD",
                "provided_value": req.EndDate,
            })
    }

    daysHeld := int(endDate.Sub(startDate).Hours() / 24)
    if daysHeld < 0 {
        return SendError(c, fiber.StatusBadRequest, ErrCodeValidationFailed, 
            "End date must be after start date", map[string]string{
                "start_date": req.StartDate,
                "end_date": req.EndDate,
            })
    }

    collection := h.db.Collection(c.Params("year"))
    var stock bson.M
    err = collection.FindOne(c.Context(), bson.M{"code": req.StockCode}).Decode(&stock)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Stock not found"})
    }

    purificationStr := safeString(stock["purification"])
    purificationRate, err := strconv.ParseFloat(purificationStr, 64)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Invalid purification rate in database"})
    }

    purificationAmount := float64(req.NumberOfStocks) * purificationRate * float64(daysHeld) / 365.0

    response := PurificationResponse{
        PurificationAmount: purificationAmount,
        DaysHeld:          daysHeld,
        PurificationRate:  purificationRate,
    }

    return c.JSON(response)
}