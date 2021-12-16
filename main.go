package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	Er      = godotenv.Load()
	Secret  = os.Getenv("PAYSTACK_SECRET_KEY")
	BaseURL = os.Getenv("PAYSTACK_BASE_URL")
)

func main() {
	if Er != nil {
		log.Fatalf("error loading .env file (%s)", Er.Error())
	}

	myRouter := gin.Default()
	myRouter.POST("/verify", func(c *gin.Context) {

		var payload struct{}
		if er := c.ShouldBindJSON(&payload); er != nil {
			c.JSONP(http.StatusBadRequest, map[string]interface{}{
				"status":  false,
				"message": "bad request",
				"error":   er.Error(),
			})
			return
		}

		//c.Writer.Header().Set("Authorization", "Bearer "+Secret)

	})
}
