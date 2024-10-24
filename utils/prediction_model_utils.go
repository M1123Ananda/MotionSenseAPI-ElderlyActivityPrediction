package utils

import (
	"MotionSense/internal/models"
	"fmt"
)

func PrepareInputData(req models.PredictActivityRequest) ([]float32, error) {
	if len(req.Acceleration) != 128 || len(req.AngularVelocity) != 128 {
		return nil, fmt.Errorf("invalid orientation data, expected 128 Axis values")
	}

	flattenedInput := make([]float32, 0, 128*6)

	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.Acceleration[i].X)
	}
	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.Acceleration[i].Y)
	}
	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.Acceleration[i].Z)
	}
	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.AngularVelocity[i].X)
	}
	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.AngularVelocity[i].Y)
	}
	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.AngularVelocity[i].Z)
	}

	return flattenedInput, nil
}

func Argmax(values []float32) int {
	maxIndex := 0
	maxValue := values[0]
	for i, value := range values {
		if value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}
	return maxIndex
}

func GetLabel(index int) string {
	labels := map[int]string{
		0: "WALKING",
		1: "WALKING_UPSTAIRS",
		2: "WALKING_DOWNSTAIRS",
		3: "SITTING",
		4: "STANDING",
		5: "LAYING",
	}
	return labels[index]
}
