package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
	"paystack/model"
	"paystack/pkg/utils"
)

var (
	Secret  = os.Getenv("PAYSTACK_SECRET_KEY")
	BaseURL = os.Getenv("PAYSTACK_BASE_URL")
)

func VerifyTransaction(c *gin.Context) {

	var payload struct {
		Reference string `json:"reference" binding:"required"`
	}

	if er := c.ShouldBindJSON(&payload); er != nil {
		c.JSONP(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "bad request",
			"error":   er.Error(),
		})
		return
	}

	client := resty.New()
	resp, er := client.R().
		SetHeader("Authorization", "Bearer "+Secret).
		Get(utils.ResolveURL(BaseURL, "/transaction/verify/"))

	if er != nil {
		c.JSONP(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  false,
			"message": "unable to reach paystack",
			"error":   er.Error(),
		})
		return
	}

	if resp.StatusCode() != 200 {
		c.JSONP(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  false,
			"message": "error confirming transaction",
			"error":   resp.String(),
		})
		return
	}

	var wResponse model.WebHookResponse
	if er := json.Unmarshal(resp.Body(), &wResponse); er != nil {
		c.JSONP(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  false,
			"message": "error confirming transaction",
			"error":   er.Error(),
		})
		return
	}

	if wResponse.Data.Status != "success" {
		c.JSONP(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "transaction not successful",
			"data":    wResponse,
		})
		return
	}

	c.JSONP(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "transaction successful",
		"data":    wResponse,
	})
}
