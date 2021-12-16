package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"paystack/model"
)

func WebHook(c *gin.Context) {

	var payload model.WebHookResponse
	if er := c.ShouldBindJSON(&payload); er != nil {
		c.JSONP(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "bad request",
			"error":   er.Error(),
		})
		return
	}

	if payload.Data.Status != "success" {
		c.JSONP(http.StatusOK, map[string]interface{}{
			"status":  payload.Data.Status,
			"message": "transaction not successful",
			"data":    payload.Data,
		})
		return
	}

	log.Println(payload.Data)

	c.JSONP(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "transaction successful",
		"data":    payload.Data,
	})
}
