package main

import (
	config "MotionSense/configs"
	"MotionSense/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	const port = "6969"
	config.LoadConfig("configs/config.yaml")

	router := gin.Default()
	router.POST("/predict", service.PredictActivity)

	if err := router.Run("0.0.0.0:" + port); err != nil {
		fmt.Println(fmt.Errorf("failed to run router on port :%s", port))
	}
}
