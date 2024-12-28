package handlers

// ###############################################################################
// Stock Handler Functions
// ###############################################################################

import (
	"strconv"
	"time"

	"github.com/anqorithm/naqa-api/internal/constants"
	"github.com/anqorithm/naqa-api/internal/models"
	"github.com/anqorithm/naqa-api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetStocksHandler fetches all stocks for a given year
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
	result := make([]models.Stock, 0, len(stocks))
	for _, doc := range stocks {
		id, _ := doc["_id"].(primitive.ObjectID)
		stock := models.Stock{
			ID:            id,
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

// GetStockByCodeHandler fetches a stock by its code for a given year
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
	result := make([]models.Stock, 0, len(stocks))
	for _, doc := range stocks {
		id, _ := doc["_id"].(primitive.ObjectID)
		stock := models.Stock{
			ID:            id,
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

// CalculatePurificationHandler calculates purification amount for a stock
func (h *Handler) CalculatePurificationHandler(c *fiber.Ctx) error {
	const daysInYear = 365.0

	var req models.PurificationRequest
	if err := c.BodyParser(&req); err != nil {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeInvalidRequest,
			constants.MsgInvalidRequestBody, nil)
	}

	if errors := utils.ValidateRequest(&req); len(errors) > 0 {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeValidationFailed,
			constants.MsgValidationFailed, fiber.Map{"errors": errors})
	}

	startDate, err := time.Parse(constants.DateFormat, req.StartDate)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeInvalidRequest,
			constants.MsgInvalidStartDateFormat, nil)
	}

	endDate, err := time.Parse(constants.DateFormat, req.EndDate)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeInvalidRequest,
			constants.MsgInvalidEndDateFormat, nil)
	}

	if !endDate.After(startDate) {
		return SendError(c, fiber.StatusBadRequest, models.ErrCodeValidationFailed,
			constants.MsgEndDateAfterStartDate, nil)
	}

	daysHeld := int(endDate.Sub(startDate).Hours() / 24)

	collection := h.db.Collection(c.Params("year"))
	var stock bson.M
	err = collection.FindOne(c.Context(), bson.M{"code": req.StockCode}).Decode(&stock)
	if err != nil {
		return SendError(c, fiber.StatusNotFound, models.ErrCodeNotFound,
			constants.MsgStockNotFound, nil)
	}

	id, _ := stock["_id"].(primitive.ObjectID)
	purificationStr := utils.SafeString(stock["purification"])
	purificationRate, err := strconv.ParseFloat(purificationStr, 64)
	if err != nil {
		return SendError(c, fiber.StatusInternalServerError, models.ErrCodeInvalidData,
			constants.MsgInvalidPurificationRate, nil)
	}

	purificationAmount := float64(req.NumberOfStocks) * purificationRate * float64(daysHeld) / daysInYear

	response := models.PurificationResponse{
		ID:                 id,
		PurificationAmount: purificationAmount,
		DaysHeld:           daysHeld,
		PurificationRate:   purificationRate,
	}

	return c.JSON(response)
}