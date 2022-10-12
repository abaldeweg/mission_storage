package main

import (
	"os"

	"github.com/abaldeweg/mission_storage/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("./.env"); err == nil {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}

	gin.SetMode(getGinMode())

	router.Router()
}

func getGinMode() string {
	switch os.Getenv("ENV") {
	case "prod":
		return gin.ReleaseMode
	case "dev":
		return gin.DebugMode
	case "test":
		return gin.TestMode
	default:
		return gin.ReleaseMode
	}
}
