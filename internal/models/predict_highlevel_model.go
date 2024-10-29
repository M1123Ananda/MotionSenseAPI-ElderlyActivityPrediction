package models

type PredictHighLevelRequest struct {
	Predictions []string `json:"predictions"`
	Positions   []string `json:"positions"`
	Time        []string `json:"time"`
}

type PredictHighLevelResponse struct {
	HighLevelActivities []string `json:"high_level_activities"`
	HighLevelTime       []string `json:"high_level_time"`
}
