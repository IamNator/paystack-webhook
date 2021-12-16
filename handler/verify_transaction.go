package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"paystack/model"
	"paystack/pkg/utils"
)

var (
	_       = godotenv.Load()
	Secret  = os.Getenv("PAYSTACK_SECRET_KEY")
	BaseURL = os.Getenv("PAYSTACK_BASE_URL")
)

func CheckTransaction(c *gin.Context) {

	var payload struct {
		Reference string `form:"reference" binding:"required"`
	}

	if er := c.ShouldBindQuery(&payload); er != nil {
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
		Get(utils.ResolveURL(BaseURL, fmt.Sprintf("/transaction/verify/%s", payload.Reference)))

	if er != nil {
		c.JSONP(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  false,
			"message": "unable to reach paystack",
			"error":   er.Error(),
		})
		return
	}

	var ErrResp model.WebHookResponseError
	if resp.StatusCode() != 200 {
		json.Unmarshal(resp.Body(), &ErrResp)
		c.JSONP(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  false,
			"message": ErrResp.Message,
		})
		return
	}

	var wResponse model.WebHookResponse
	if er := json.Unmarshal(resp.Body(), &wResponse); er != nil {
		json.Unmarshal(resp.Body(), &ErrResp)
		c.JSONP(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  false,
			"message": ErrResp.Message,
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
