package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Configuration *Config

type Config struct {
	ModelName               string `yaml:"model_name"`
	TorchPredictionEndpoint string `yaml:"torch_prediction_endpoint"`
}

func LoadConfig(path string) {
	Configuration = &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, Configuration)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("Successfully Loaded Configs")
}
