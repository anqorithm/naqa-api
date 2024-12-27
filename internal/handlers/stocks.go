package handlers
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func getStringValue(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return ""
	}
}

func (h *Handler) GetStocksByYearHandler(c *fiber.Ctx) error {
	// ...existing code...
	for _, doc := range stockDocs {
		stock := Stock{
			Name:          getStringValue(doc["name"]),
			Code:          getStringValue(doc["code"]),
			Sector:        getStringValue(doc["sector"]),
			ShariaOpinion: getStringValue(doc["sharia_opinion"]),
		}
		// ...existing code...
	}
	// ...existing code...
}

func (h *Handler) SearchStocksHandler(c *fiber.Ctx) error {
	// ...existing code...
	for _, doc := range stockDocs {
		stock := Stock{
			Name:          getStringValue(doc["name"]),
			Code:          getStringValue(doc["code"]),
			Sector:        getStringValue(doc["sector"]),
			ShariaOpinion: getStringValue(doc["sharia_opinion"]),
		}
		// ...existing code...
	}
	// ...existing code...
}

import (
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Stock represents a single stock item
type Stock struct {
	Name          string `json:"name" bson:"name" example:"Saudi Aramco"`
	Code          string `json:"code" bson:"code" example:"2222"`
	Sector        string `json:"sector" bson:"sector" example:"Energy"`
	ShariaOpinion string `json:"sharia_opinion" bson:"sharia_opinion" example:"compliant"`
}

// StockResponse represents the response structure
// swagger:model
type StockResponse struct {
	// List of stocks
	// swagger:allOf
	Stocks []Stock `json:"stocks"`
}

// ErrorResponse represents the error response
type ErrorResponse struct {
	Error string `json:"error" example:"Failed to fetch stocks"`
}

// swagger:response stocksResponse
type swaggerStocksResponse struct {
	// in: body
	Body StockResponse
}

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

// @Summary Get stocks by year
// @Description Retrieves all stocks for a specific year from the database
// @Tags stocks
// @Accept json
// @Produce json
// @Param year path string true "Year of stocks data" example:"2023"
// @Success 200 {object} StockResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /stocks/year/{year} [get]
func (h *Handler) GetStocksByYearHandler(c *fiber.Ctx) error {
	year := c.Params("year")
	
	collection := h.db.Collection(year)
	fmt.Println(collection)
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(ErrorResponse{Error: "Failed to fetch stocks"})
	}

	var stockDocs []bson.M
	if err = cursor.All(c.Context(), &stockDocs); err != nil {
		return c.Status(500).JSON(ErrorResponse{Error: "Failed to parse stocks"})
	}

	// Convert bson.M to []Stock
	stocks := make([]Stock, 0, len(stockDocs))
	for _, doc := range stockDocs {
		stock := Stock{
			Name:          doc["name"].(string),
			Code:          doc["code"].(string),
			Sector:        doc["sector"].(string),
			ShariaOpinion: doc["sharia_opinion"].(string),
		}
		stocks = append(stocks, stock)
	}
	
	return c.JSON(StockResponse{Stocks: stocks})
}

// @Summary Search stocks
// @Description Search stocks with various filters
// @Tags stocks
// @Accept json
// @Produce json
// @Param year path string true "Year of stocks" example:"2023"
// @Param name query string false "Stock name" example:"Aramco"
// @Param code query string false "Stock code" example:"2222"
// @Param sector query string false "Stock sector" example:"Energy"
// @Param sharia_opinion query string false "Sharia compliance status" example:"compliant"
// @Success 200 {object} StockResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /stocks/year/{year}/search [get]
func (h *Handler) SearchStocksHandler(c *fiber.Ctx) error {
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
        return c.Status(500).JSON(ErrorResponse{Error: "Failed to search stocks"})
    }
    
    var stockDocs []bson.M
    if err = cursor.All(c.Context(), &stockDocs); err != nil {
        return c.Status(500).JSON(ErrorResponse{Error: "Failed to parse stocks"})
    }

    // Convert bson.M to []Stock
    stocks := make([]Stock, 0, len(stockDocs))
    for _, doc := range stockDocs {
        stock := Stock{
            Name:          doc["name"].(string),
            Code:          doc["code"].(string),
            Sector:        doc["sector"].(string),
            ShariaOpinion: doc["sharia_opinion"].(string),
        }
        stocks = append(stocks, stock)
    }
    return c.JSON(StockResponse{Stocks: stocks})}