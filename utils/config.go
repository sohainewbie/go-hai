package utils

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

var Config config

type config struct {
	Service struct {
		Name    string `json:"name"`
		Port    string `json:"port"`
		BaseURL string `json:"baseUrl"`
		Version string `json:"version"`
	}
	Database struct {
		Mysql struct {
			Name     string `json:"name"`
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			LogMode  bool   `json:"logMode"`
		}`json:"mysql"`
		Postgres struct {
			Name     string `json:"name"`
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			SslMode  string `json:"ssl_mode"`
			LogMode  bool   `json:"logMode"`
		}`json:"postgres"`		
		Redis struct {
			DB       int `json:"db"`
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Password string `json:"password"`
		}`json:"redis"`
	}
	InternalService struct {
	} `json:"internalService"`
	DebugMode  bool   `json:"debugMode"`
	SecretKey  string `json:"secretKey"`
}

func init() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(file), &Config)
}
