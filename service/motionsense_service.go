package service

import (
	"MotionSense/internal/models"
	"MotionSense/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PredictActivity(context *gin.Context) {
	var req models.PredictActivityRequest

	if err := context.BindJSON(&req); err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	model, backend, err := utils.LoadModel("dummy_model.onnx")
	if err != nil {
		fmt.Println("failed to load model: " + err.Error())
	}

	fmt.Println(model)

	inputTensor, err := utils.PrepareInputTensor(req)
	if err != nil {
		fmt.Println("failed to prepare input tensor: " + err.Error())
		return
	}

	if err := model.SetInput(0, inputTensor); err != nil {
		fmt.Println("failed to input to model: " + err.Error())
	}

	if err := backend.Run(); err != nil {
		fmt.Println("failed to run model: " + err.Error())
	}

	output, err := model.GetOutputTensors()
	if err != nil {
		fmt.Println("failed to get model output: " + err.Error())
	}

	//fmt.Println(output[0].Shape())
	//
	//for i, input := range model.GetInputTensors() {
	//	fmt.Printf("Input %d: %v\n", i, input)
	//}

	//op, err := model.GetOutputTensors()

	//for i, output := range op {
	//	fmt.Printf("Output %d: %v\n", i, output)
	//}

	context.IndentedJSON(http.StatusOK, models.PredictActivityResponse{TagNumber: req.TagNumber,
		Activity: output[0].String()})
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
