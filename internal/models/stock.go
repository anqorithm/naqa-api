package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stock struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Code          string             `json:"code" bson:"code"`
	Sector        string             `json:"sector" bson:"sector"`
	ShariaOpinion string             `json:"sharia_opinion" bson:"sharia_opinion"`
	Purification  string             `json:"purification" bson:"purification"`
}

type StockResponse struct {
	Stocks []Stock `json:"stocks"`
}


type PurificationRequest struct {
	StartDate      string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate        string `json:"end_date" validate:"required,datetime=2006-01-02"`
	NumberOfStocks int    `json:"number_of_stocks" validate:"required,gt=0"`
	StockCode      string `json:"stock_code" validate:"required"`
}


type PurificationResponse struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PurificationAmount float64            `json:"purification_amount"`
	DaysHeld           int                `json:"days_held"`
	PurificationRate   float64            `json:"purification_rate"`
}
