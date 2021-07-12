package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetString(key, defaultValue string) string {
 	data := os.Getenv(key)
	if data != "" {
		return data
	}
	return defaultValue
}

func GetInt(key string,defaultValue int) int {
	b,err := strconv.Atoi(GetString(key,""))
	if err != nil {
		return defaultValue
	}
	return b
}
