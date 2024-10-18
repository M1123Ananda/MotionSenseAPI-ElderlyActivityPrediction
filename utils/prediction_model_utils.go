package utils

import (
	"MotionSense/internal/models"
	"fmt"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
	"os"
)

func LoadModel(modelName string) (m *onnx.Model, be *gorgonnx.Graph, err error) {
	backend := gorgonnx.NewGraph()
	model := onnx.NewModel(backend)

	b, err := os.ReadFile("ML/" + modelName)
	if err != nil {
		return nil, nil, err
	}

	if err := model.UnmarshalBinary(b); err != nil {
		return nil, nil, err
	}

	return model, backend, nil
}

func PrepareInputTensor(req models.PredictActivityRequest) (tensor.Tensor, error) {
	if len(req.Acceleration) != 128 || len(req.AngularVelocity) != 128 {
		return nil, fmt.Errorf("invalid orientation data, expected 128 Axis values")
	}

	flattenedInput := make([]float32, 0, 128*6)

	for i := 0; i < 128; i++ {
		flattenedInput = append(flattenedInput, req.Acceleration[i].X, req.Acceleration[i].Y, req.Acceleration[i].Z)
		flattenedInput = append(flattenedInput, req.AngularVelocity[i].X, req.AngularVelocity[i].Y, req.AngularVelocity[i].Z)
	}

	t := tensor.New(tensor.WithShape(128, 6), tensor.Of(tensor.Float32), tensor.WithBacking(flattenedInput))

	return t, nil
}
