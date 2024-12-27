package handlers

import (
	"strconv"
	"time"

	"github.com/anqorithm/naqa-api/internal/models"
	"github.com/anqorithm/naqa-api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// ###############################################################################
// Stock Handler Functions
// ###############################################################################

func (h *Handler) GetStocksByYearHandler(c *fiber.Ctx) error {
	collection := h.db.Collection(c.Params("year"))

	var stocks []bson.M
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return SendError(c, fiber.StatusInternalServerError, models.ErrCodeDatabaseError,
			"Failed to fetch stocks", nil)
	}
	defer cursor.Close(c.Context())

	if err := cursor.All(c.Context(), &stocks); err != nil {
		return SendError(c, fiber.StatusInternalServerError, models.ErrCodeDatabaseError,
			"Failed to parse stocks", nil)
	}
	if len(stocks)==0 {
		return c.Status(404).JSON(fiber.Map{"error": "Stock Not Found We could not find any stock matching your criteria. Please verify your search parameters and try again. "});
	}

	result := make([]models.Stock, 0, len(stocks))
	for _, doc := range stocks {
		stock := models.Stock{
			Name:          utils.SafeString(doc["name"]),
			Code:          utils.SafeString(doc["code"]),
			Sector:        utils.SafeString(doc["sector"]),
			ShariaOpinion: utils.SafeString(doc["sharia_opinion"]),
			Purification:  utils.SafeString(doc["purification"]),
		}
		result = append(result, stock)
	}

	return c.JSON(models.StockResponse{Stocks: result})
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
	if len(stocks)==0 {
		return c.Status(404).JSON(fiber.Map{"error": "Stock Not Found We could not find any stock matching your criteria. Please verify your search parameters and try again. "});
	}
	result := make([]models.Stock, 0, len(stocks))
	for _, doc := range stocks {
		stock := models.Stock{
			Name:          utils.SafeString(doc["name"]),
			Code:          utils.SafeString(doc["code"]),
			Sector:        utils.SafeString(doc["sector"]),
			ShariaOpinion: utils.SafeString(doc["sharia_opinion"]),
			Purification:  utils.SafeString(doc["purification"]),
		}
		result = append(result, stock)
	}

	return c.JSON(models.StockResponse{Stocks: result})
}

func (h *Handler) CalculatePurificationHandler(c *fiber.Ctx) error {
	var req models.PurificationRequest
	if err := c.BodyParser(&req); err != nil {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeInvalidRequest,
			"Invalid request body", nil)
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeInvalidDateFormat,
			"Invalid start date format", map[string]string{
				"expected_format": "YYYY-MM-DD",
				"provided_value":  req.StartDate,
			})
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeInvalidDateFormat,
			"Invalid end date format", map[string]string{
				"expected_format": "YYYY-MM-DD",
				"provided_value":  req.EndDate,
			})
	}

	daysHeld := int(endDate.Sub(startDate).Hours() / 24)
	if daysHeld < 0 {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeValidationFailed,
			"End date must be after start date", map[string]string{
				"start_date": req.StartDate,
				"end_date":   req.EndDate,
			})
	}

	collection := h.db.Collection(c.Params("year"))
	var stock bson.M
	err = collection.FindOne(c.Context(), bson.M{"code": req.StockCode}).Decode(&stock)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Stock not found"})
	}

	purificationStr := utils.SafeString(stock["purification"])
	purificationRate, err := strconv.ParseFloat(purificationStr, 64)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid purification rate in database"})
	}

	purificationAmount := float64(req.NumberOfStocks) * purificationRate * float64(daysHeld) / 365.0

	response := models.PurificationResponse{
		PurificationAmount: purificationAmount,
		DaysHeld:           daysHeld,
		PurificationRate:   purificationRate,
	}

	return c.JSON(response)
}
