package service

import (
	config "MotionSense/configs"
	"MotionSense/internal/models"
	"MotionSense/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func PredictActivity(context *gin.Context) {
	var req models.PredictActivityRequest

	if err := context.BindJSON(&req); err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		context.IndentedJSON(http.StatusInternalServerError, models.PredictActivityResponse{TagNumber: req.TagNumber,
			Activity: ""})
		return
	}

	input, err := utils.PrepareInputData(req)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		context.IndentedJSON(http.StatusInternalServerError, models.PredictActivityResponse{TagNumber: req.TagNumber,
			Activity: ""})
		return
	}

	body, err := json.Marshal(map[string]any{
		"data": input,
	})

	resp, err := utils.CallPostRequest(config.Configuration.TorchPredictionEndpoint, body)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		context.IndentedJSON(http.StatusInternalServerError, models.PredictActivityResponse{TagNumber: req.TagNumber,
			Activity: ""})
		return
	}
	defer resp.Body.Close()

	var respBodyModel models.TorchPredictionResponse
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		context.IndentedJSON(http.StatusInternalServerError, models.PredictActivityResponse{TagNumber: req.TagNumber,
			Activity: ""})
		return
	}

	err = json.Unmarshal(responseBody, &respBodyModel)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		context.IndentedJSON(http.StatusInternalServerError, models.PredictActivityResponse{TagNumber: req.TagNumber,
			Activity: ""})
		return
	}

	if resp.StatusCode != http.StatusOK {
		context.IndentedJSON(resp.StatusCode, models.PredictActivityResponse{TagNumber: req.TagNumber,
			Activity: ""})
		return
	}

	maxIndex := utils.Argmax(respBodyModel.Prediction)
	labeledPrediction := utils.GetLabel(maxIndex)

	context.IndentedJSON(http.StatusOK, models.PredictActivityResponse{TagNumber: req.TagNumber,
		Activity: labeledPrediction})
	return
}

func PredictHighLevel(context *gin.Context) {
	const modeFilterWindow = 3
	var req models.PredictHighLevelRequest

	if err := context.BindJSON(&req); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, models.PredictHighLevelResponse{
			HighLevelActivities: nil,
			HighLevelTime:       nil,
		})
		return
	}

	if len(req.Predictions) != len(req.Positions) {
		fmt.Println(fmt.Errorf("predictions and positions array length is not equal %d, %d", len(req.Predictions), len(req.Positions)))
		context.IndentedJSON(http.StatusInternalServerError, models.PredictHighLevelResponse{
			HighLevelActivities: nil,
			HighLevelTime:       nil,
		})
		return
	}

	cleanPredictions := utils.ModeFilter(&req.Predictions, 5)
	cleanPositions := utils.ModeFilter(&req.Positions, 5)
	groupedPredictions, groupedPositions, groupedTime := utils.GroupData(cleanPredictions, cleanPositions, &req.Time)
	highLevelActivities, highLevelTime := utils.CombineAndTransform(groupedPredictions, groupedPositions, groupedTime)

	context.IndentedJSON(http.StatusOK, models.PredictHighLevelResponse{
		HighLevelActivities: *highLevelActivities,
		HighLevelTime:       *highLevelTime,
	})
	return
}
