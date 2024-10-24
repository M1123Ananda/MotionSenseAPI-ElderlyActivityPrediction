package models

type Axis struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type PredictActivityRequest struct {
	TagNumber       int    `json:"tag_number"`
	Acceleration    []Axis `json:"acceleration"`
	AngularVelocity []Axis `json:"angular_velocity"`
}

type PredictActivityResponse struct {
	TagNumber int    `json:"tag_number"`
	Activity  string `json:"activity"`
}

type TorchPredictionResponse struct {
	Prediction []float32 `json:"prediction"`
	Message    string    `json:"message"`
}
