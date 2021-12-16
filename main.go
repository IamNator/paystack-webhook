package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"paystack/handler"
)

var (
	Er = godotenv.Load()
)

func main() {
	if Er != nil {
		log.Fatalf("error loading .env file (%s)", Er.Error())
	}

	myRouter := gin.Default()
	myRouter.POST("/verify", handler.VerifyTransaction)
}
