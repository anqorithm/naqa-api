package models

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
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	NumberOfStocks int    `json:"number_of_stocks"`
	StockCode      string `json:"stock_code"`
}

type PurificationResponse struct {
	PurificationAmount float64 `json:"purification_amount"`
	DaysHeld           int     `json:"days_held"`
	PurificationRate   float64 `json:"purification_rate"`
}