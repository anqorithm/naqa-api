package models

// ###############################################################################
// Stock Model
// ###############################################################################

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Stock represents a stock entity
type Stock struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Code          string             `json:"code" bson:"code" example:"1111"`
	Sector        string             `json:"sector" bson:"sector" example:"الطاقة"`
	ShariaOpinion string             `json:"sharia_opinion" bson:"sharia_opinion" example:"نقية"`
	Purification  string             `json:"purification" bson:"purification" example:"0.25"`
	NameAr        string             `json:"name_ar" bson:"name_ar" example:"شركة الطاقة السعودية"`
	NameEn        string             `json:"name_en" bson:"name_en" example:"Saudi Energy Company"`
	Logo          string             `json:"logo" bson:"logo" example:"https://example.com/logo.png"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at" swaggertype:"string" format:"date-time"`
}

// StockResponse represents a response containing a list of stocks
type StockResponse struct {
	Stocks []Stock `json:"stocks"`
}

// PurificationRequest represents a request to calculate purification amount
type PurificationRequest struct {
	StartDate      string `json:"start_date" validate:"required,datetime=2006-01-02" example:"2023-01-01"`
	EndDate        string `json:"end_date" validate:"required,datetime=2006-01-02" example:"2023-12-31"`
	NumberOfStocks int    `json:"number_of_stocks" validate:"required,gt=0" example:"100"`
	StockCode      string `json:"stock_code" validate:"required" example:"1111"`
}

// PurificationResponse represents a response containing the purification amount
type PurificationResponse struct {
	ID                 primitive.ObjectID       `json:"_id" bson:"_id,omitempty"`
	PurificationAmount float64                  `json:"purification_amount" example:"25.5"`
	DaysHeld           int                      `json:"days_held" example:"365"`
	PurificationRate   float64                  `json:"purification_rate" example:"0.25"`
	YearlyBreakdown    []YearlyPurificationInfo `json:"yearly_breakdown,omitempty"`
	IsMultiYear        bool                     `json:"is_multi_year" example:"false"`
	Warnings           []string                 `json:"warnings,omitempty"`
}

// YearlyPurificationInfo represents purification calculation for a specific year
type YearlyPurificationInfo struct {
	Year               int     `json:"year"`
	PurificationRate   float64 `json:"purification_rate"`
	DaysInPeriod       int     `json:"days_in_period"`
	TotalDaysInYear    int     `json:"total_days_in_year"`
	YearProportion     float64 `json:"year_proportion"`
	PurificationAmount float64 `json:"purification_amount"`
	CompanyNameEn      string  `json:"company_name_en"`
	CompanyNameAr      string  `json:"company_name_ar"`
	ShariaOpinion      string  `json:"sharia_opinion"`
}

// ###############################################################################
// API Response Models
// ###############################################################################

// RootResponse represents the root API endpoint response
type RootResponse struct {
	AppName      string   `json:"app_name" example:"Naqa API"`
	Version      string   `json:"version" example:"1.0.0"`
	Description  string   `json:"description" example:"A stock analysis and Sharia compliance API"`
	RequestTime  string   `json:"request_time" example:"2023-01-01T00:00:00Z"`
	Environment  string   `json:"environment" example:"development"`
	SupportedAPI []string `json:"supported_api" example:"/api/v1"`
}

// ApiV1Response represents the API v1 endpoint response
type ApiV1Response struct {
	Status     string   `json:"status" example:"active"`
	Message    string   `json:"message" example:"Welcome to NAQA API v1"`
	Version    string   `json:"version" example:"1.0.0"`
	Env        string   `json:"env" example:"development"`
	ServerTime string   `json:"server_time" example:"2023-01-01T00:00:00Z"`
	RequestID  string   `json:"request_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Endpoints  []string `json:"endpoints" example:"/api/v1/stocks,/api/v1/health,/api/v1/metrics"`
}

// HealthResponse represents the health check endpoint response
type HealthResponse struct {
	Status    string `json:"status" example:"OK"`
	Timestamp string `json:"timestamp" example:"2023-01-01T00:00:00Z"`
	Message   string `json:"message" example:"API is healthy"`
	RequestID string `json:"request_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// EndpointInfo represents information about an API endpoint
type EndpointInfo struct {
	Name        string                 `json:"name" example:"Get Stocks"`
	Path        string                 `json:"path" example:"/api/v1/stocks/year/{year}"`
	Method      string                 `json:"method" example:"GET"`
	Description string                 `json:"description" example:"Get all stocks for a specific year"`
	Example     interface{}            `json:"example,omitempty"`
	Parameters  []string               `json:"parameters,omitempty" example:"name,code,sector"`
}

// StocksBaseResponse represents the base stocks endpoint response
type StocksBaseResponse struct {
	Status          string         `json:"status" example:"success"`
	Message         string         `json:"message" example:"Welcome to Stocks API"`
	AvailableYears  []string       `json:"available_years" example:"2023,2024"`
	Endpoints       []EndpointInfo `json:"endpoints"`
	Documentation   string         `json:"documentation" example:"https://github.com/anqorithm/naqa-api"`
}