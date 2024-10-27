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

	fmt.Println(input)
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
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	if len(req.Predictions) != len(req.Positions) {
		fmt.Println(fmt.Errorf("predictions and positions array length is not equal"))
		return
	}

	//clean_predictions := utils.ModeFilter(req.Predictions, 5)
	//clean_positions := utils.ModeFilter(req.Positions, 5)
}
