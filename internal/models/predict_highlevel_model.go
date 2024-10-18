package models

type PredictHighLevelRequest struct {
	Predictions []string `json:"predictions"`
	Positions   []string `json:"position"`
}

type PredictHighLevelResponse struct {
	HighLevel []string `json:"high_level"`
}
