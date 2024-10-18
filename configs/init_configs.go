package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Configuration *Config

//var DBConfiguration *DBConfig

type Config struct {
	ModelName string `yaml:"model_name"`
}

//type DBConfig struct {
//	DBUser     string
//	DBPassword string
//	DBName     string
//	DBHost     string
//	DBPort     string
//}

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

	//DBConfiguration = &DBConfig{
	//	DBUser:     os.Getenv("DB_USER"),
	//	DBPassword: os.Getenv("DB_PASSWORD"),
	//	DBName:     os.Getenv("DB_NAME"),
	//	DBHost:     os.Getenv("DB_HOST"),
	//	DBPort:     os.Getenv("DB_PORT"),
	//}

	log.Println("Successfully Loaded Configs")
}

//func (c *DBConfig) GetDSN() string {
//	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
//		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
//}
