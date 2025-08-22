package handlers

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

// GetStocksByYearHandler handles getting stocks for a specific year
// @Summary Get all stocks for a specific year
// @Description Retrieve all stocks and their information for a given year
// @Tags Stocks
// @Accept json
// @Produce json
// @Param year path string true "Year (e.g., 2024)"
// @Success 200 {object} models.StockResponse "List of stocks"
// @Failure 400 {object} map[string]interface{} "Invalid year"
// @Failure 500 {object} map[string]interface{} "Database error"
// @Router /api/v1/stocks/year/{year} [get]
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
			Code:          utils.SafeString(doc["code"]),
			Sector:        utils.SafeString(doc["sector"]),
			ShariaOpinion: utils.SafeString(doc["sharia_opinion"]),
			Purification:  utils.SafeString(doc["purification"]),
			NameAr:        utils.SafeString(doc["name_ar"]),
			NameEn:        utils.SafeString(doc["name_en"]),
			Logo:          utils.SafeString(doc["logo"]),
		}
		
		if createdAt, ok := doc["created_at"].(primitive.DateTime); ok {
			stock.CreatedAt = createdAt.Time()
		}
		if updatedAt, ok := doc["updated_at"].(primitive.DateTime); ok {
			stock.UpdatedAt = updatedAt.Time()
		}
		
		result = append(result, stock)
	}

	return c.JSON(models.StockResponse{Stocks: result})
}

// SearchStocksHandler handles searching stocks with filters
// @Summary Search stocks with filters
// @Description Search for stocks using various filters like name, code, sector, or sharia opinion
// @Tags Stocks
// @Accept json
// @Produce json
// @Param year path string true "Year (e.g., 2024)"
// @Param name query string false "Stock name (Arabic or English)"
// @Param code query string false "Stock code"
// @Param sector query string false "Business sector"
// @Param sharia_opinion query string false "Sharia compliance opinion"
// @Param purification query string false "Purification percentage"
// @Success 200 {object} models.StockResponse "Filtered list of stocks"
// @Failure 400 {object} map[string]interface{} "Invalid parameters"
// @Failure 500 {object} map[string]interface{} "Database error"
// @Router /api/v1/stocks/year/{year}/search [get]
func (h *Handler) SearchStocksHandler(c *fiber.Ctx) error {
	filter := bson.M{}

	if name := c.Query("name"); name != "" {
		filter["$or"] = []bson.M{
			{"name_ar": bson.M{"$regex": name, "$options": "i"}},
			{"name_en": bson.M{"$regex": name, "$options": "i"}},
		}
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
			Code:          utils.SafeString(doc["code"]),
			Sector:        utils.SafeString(doc["sector"]),
			ShariaOpinion: utils.SafeString(doc["sharia_opinion"]),
			Purification:  utils.SafeString(doc["purification"]),
			NameAr:        utils.SafeString(doc["name_ar"]),
			NameEn:        utils.SafeString(doc["name_en"]),
			Logo:          utils.SafeString(doc["logo"]),
		}
		
		if createdAt, ok := doc["created_at"].(primitive.DateTime); ok {
			stock.CreatedAt = createdAt.Time()
		}
		if updatedAt, ok := doc["updated_at"].(primitive.DateTime); ok {
			stock.UpdatedAt = updatedAt.Time()
		}
		
		result = append(result, stock)
	}

	return c.JSON(models.StockResponse{Stocks: result})
}

// CalculatePurificationHandler handles purification calculation requests
// @Summary Calculate stock purification amount
// @Description Calculate the purification amount for a stock based on holding period and number of shares
// @Tags Stocks
// @Accept json
// @Produce json
// @Param year path string true "Year (e.g., 2024)"
// @Param request body models.PurificationRequest true "Purification calculation request"
// @Success 200 {object} models.PurificationResponse "Purification calculation result"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Stock not found"
// @Failure 500 {object} map[string]interface{} "Server error"
// @Router /api/v1/stocks/year/{year}/calculate-purification [post]
func (h *Handler) CalculatePurificationHandler(c *fiber.Ctx) error {
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

	response, err := h.calculateMultiYearPurification(c, req, startDate, endDate)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (h *Handler) calculateMultiYearPurification(c *fiber.Ctx, req models.PurificationRequest, startDate, endDate time.Time) (*models.PurificationResponse, error) {
	years := GetYearsInPeriod(startDate, endDate)
	
	var totalPurification float64
	var yearlyBreakdown []models.YearlyPurificationInfo
	var warnings []string
	var totalDaysHeld int
	var lastFoundID primitive.ObjectID

	isMultiYear := len(years) > 1
	if isMultiYear {
		warnings = append(warnings, "Multi-year period detected. Each year calculated separately.")
	}
	for _, year := range years {
		collection := h.db.Collection(strconv.Itoa(year))
		var stock bson.M
		err := collection.FindOne(c.Context(), bson.M{"code": req.StockCode}).Decode(&stock)
		if err != nil {
			warnings = append(warnings, "No data found for company "+req.StockCode+" in year "+strconv.Itoa(year))
			continue
		}

		lastFoundID, _ = stock["_id"].(primitive.ObjectID)

		purificationStr := utils.SafeString(stock["purification"])
		purificationRate, err := strconv.ParseFloat(purificationStr, 64)
		if err != nil {
			warnings = append(warnings, "Invalid purification rate for "+req.StockCode+" in year "+strconv.Itoa(year))
			continue
		}

		daysInYear := GetDaysInYearForPeriod(startDate, endDate, year)
		totalDaysInYear := GetDaysInYear(year)
		
		if daysInYear == 0 {
			continue
		}

		totalDaysHeld += daysInYear
		yearProportion := float64(daysInYear) / float64(totalDaysInYear)
		yearPurification := float64(req.NumberOfStocks) * purificationRate * yearProportion
		totalPurification += yearPurification

		yearInfo := models.YearlyPurificationInfo{
			Year:               year,
			PurificationRate:   purificationRate,
			DaysInPeriod:       daysInYear,
			TotalDaysInYear:    totalDaysInYear,
			YearProportion:     yearProportion,
			PurificationAmount: yearPurification,
			CompanyNameEn:      utils.SafeString(stock["name_en"]),
			CompanyNameAr:      utils.SafeString(stock["name_ar"]),
			ShariaOpinion:      utils.SafeString(stock["sharia_opinion"]),
		}
		yearlyBreakdown = append(yearlyBreakdown, yearInfo)
	}

	if len(yearlyBreakdown) == 0 {
		return nil, SendError(c, fiber.StatusNotFound, models.ErrCodeNotFound,
			constants.MsgStockNotFound, nil)
	}

	var avgPurificationRate float64
	if len(yearlyBreakdown) > 0 {
		var totalRate float64
		for _, info := range yearlyBreakdown {
			totalRate += info.PurificationRate
		}
		avgPurificationRate = totalRate / float64(len(yearlyBreakdown))
	}

	response := &models.PurificationResponse{
		ID:                 lastFoundID,
		PurificationAmount: totalPurification,
		DaysHeld:           totalDaysHeld,
		PurificationRate:   avgPurificationRate,
		YearlyBreakdown:    yearlyBreakdown,
		IsMultiYear:        isMultiYear,
		Warnings:           warnings,
	}

	return response, nil
}

func GetYearsInPeriod(startDate, endDate time.Time) []int {
	var years []int
	currentYear := startDate.Year()
	endYear := endDate.Year()
	
	for currentYear <= endYear {
		years = append(years, currentYear)
		currentYear++
	}
	
	return years
}

func GetDaysInYearForPeriod(startDate, endDate time.Time, year int) int {
	// Get the year boundaries
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	yearEnd := time.Date(year, 12, 31, 23, 59, 59, 999999999, time.UTC)
	
	// Find the overlap between the period and the year
	overlapStart := startDate
	if yearStart.After(startDate) {
		overlapStart = yearStart
	}
	
	overlapEnd := endDate
	if yearEnd.Before(endDate) {
		overlapEnd = yearEnd
	}
	
	if overlapStart.After(overlapEnd) {
		return 0
	}
	
	return int(overlapEnd.Sub(overlapStart).Hours()/24) + 1
}

// GetDaysInYear returns the number of days in a year (considering leap years)
func GetDaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// IsLeapYear checks if a year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}