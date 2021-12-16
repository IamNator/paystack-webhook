package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	"os"
	"paystack/handler"
)

func main() {

	if Er := godotenv.Load(); Er != nil {
		log.Fatalf("error loading .env file (%s)", Er.Error())
	}

	apiLogger := zerolog.New(os.Stdout)
	myRouter := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT missing in env")
	}

	apiLogger.Info().Msgf("server running at port %s", port)
	myRouter.GET("/transaction/check", handler.CheckTransaction)

	if er := http.ListenAndServe(":"+port, myRouter); er != nil {
		log.Fatalf("unable to listen and serve at port %s", port)
	}
}
