package models

// ###############################################################################
// Stock Model
// ###############################################################################

import "go.mongodb.org/mongo-driver/bson/primitive"

// Stock represents a stock entity
type Stock struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Code          string             `json:"code" bson:"code"`
	Sector        string             `json:"sector" bson:"sector"`
	ShariaOpinion string             `json:"sharia_opinion" bson:"sharia_opinion"`
	Purification  string             `json:"purification" bson:"purification"`
	NameAr        string             `json:"name_ar" bson:"name_ar"`
	NameEn        string             `json:"name_en" bson:"name_en"`
	Logo          string             `json:"logo" bson:"logo"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

// StockResponse represents a response containing a list of stocks
type StockResponse struct {
	Stocks []Stock `json:"stocks"`
}

// PurificationRequest represents a request to calculate purification amount
type PurificationRequest struct {
	StartDate      string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate        string `json:"end_date" validate:"required,datetime=2006-01-02"`
	NumberOfStocks int    `json:"number_of_stocks" validate:"required,gt=0"`
	StockCode      string `json:"stock_code" validate:"required"`
}

// PurificationResponse represents a response containing the purification amount
type PurificationResponse struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PurificationAmount float64            `json:"purification_amount"`
	DaysHeld           int                `json:"days_held"`
	PurificationRate   float64            `json:"purification_rate"`
}