package model

import "time"

type (
	CheckTransactionErrResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	CheckTransactionResponse struct {
		Event string `json:"event"`
		Data  struct {
			Id              int         `json:"id"`
			Domain          string      `json:"domain"`
			Status          string      `json:"status"`
			Reference       string      `json:"reference"`
			Amount          int         `json:"amount"`
			Message         interface{} `json:"message"`
			GatewayResponse string      `json:"gateway_response"`
			PaidAt          time.Time   `json:"paid_at"`
			CreatedAt       time.Time   `json:"created_at"`
			Channel         string      `json:"channel"`
			Currency        string      `json:"currency"`
			IpAddress       string      `json:"ip_address"`
			Metadata        interface{} `json:"metadata"`
			Log             struct {
				TimeSpent      int           `json:"time_spent"`
				Attempts       int           `json:"attempts"`
				Authentication string        `json:"authentication"`
				Errors         int           `json:"errors"`
				Success        bool          `json:"success"`
				Mobile         bool          `json:"mobile"`
				Input          []interface{} `json:"input"`
				Channel        interface{}   `json:"channel"`
				History        []struct {
					Type    string `json:"type"`
					Message string `json:"message"`
					Time    int    `json:"time"`
				} `json:"history"`
			} `json:"log"`
			Fees     interface{} `json:"fees"`
			Customer struct {
				Id           int         `json:"id"`
				FirstName    string      `json:"first_name"`
				LastName     string      `json:"last_name"`
				Email        string      `json:"email"`
				CustomerCode string      `json:"customer_code"`
				Phone        interface{} `json:"phone"`
				//Metadata     interface{} `json:"metadata"`
				RiskAction string `json:"risk_action"`
			} `json:"customer"`
			Authorization struct {
				AuthorizationCode string `json:"authorization_code"`
				Bin               string `json:"bin"`
				Last4             string `json:"last4"`
				ExpMonth          string `json:"exp_month"`
				ExpYear           string `json:"exp_year"`
				CardType          string `json:"card_type"`
				Bank              string `json:"bank"`
				CountryCode       string `json:"country_code"`
				Brand             string `json:"brand"`
				AccountName       string `json:"account_name"`
			} `json:"authorization"`
			Plan struct {
			} `json:"plan"`
		} `json:"data"`
	}

	WebHookResponse struct {
		Event string `json:"event"`
		Data  struct {
			Id              int         `json:"id"`
			Domain          string      `json:"domain"`
			Status          string      `json:"status"`
			Reference       string      `json:"reference"`
			Amount          int         `json:"amount"`
			Message         interface{} `json:"message"`
			GatewayResponse string      `json:"gateway_response"`
			PaidAt          time.Time   `json:"paid_at"`
			CreatedAt       time.Time   `json:"created_at"`
			Channel         string      `json:"channel"`
			Currency        string      `json:"currency"`
			IpAddress       string      `json:"ip_address"`
			Metadata        interface{} `json:"metadata"`
			Log             struct {
				TimeSpent      int           `json:"time_spent"`
				Attempts       int           `json:"attempts"`
				Authentication string        `json:"authentication"`
				Errors         int           `json:"errors"`
				Success        bool          `json:"success"`
				Mobile         bool          `json:"mobile"`
				Input          []interface{} `json:"input"`
				Channel        interface{}   `json:"channel"`
				History        []struct {
					Type    string `json:"type"`
					Message string `json:"message"`
					Time    int    `json:"time"`
				} `json:"history"`
			} `json:"log"`
			Fees     interface{} `json:"fees"`
			Customer struct {
				Id           int         `json:"id"`
				FirstName    string      `json:"first_name"`
				LastName     string      `json:"last_name"`
				Email        string      `json:"email"`
				CustomerCode string      `json:"customer_code"`
				Phone        interface{} `json:"phone"`
				//Metadata     interface{} `json:"metadata"`
				RiskAction string `json:"risk_action"`
			} `json:"customer"`
			Authorization struct {
				AuthorizationCode string `json:"authorization_code"`
				Bin               string `json:"bin"`
				Last4             string `json:"last4"`
				ExpMonth          string `json:"exp_month"`
				ExpYear           string `json:"exp_year"`
				CardType          string `json:"card_type"`
				Bank              string `json:"bank"`
				CountryCode       string `json:"country_code"`
				Brand             string `json:"brand"`
				AccountName       string `json:"account_name"`
			} `json:"authorization"`
			Plan struct {
			} `json:"plan"`
		} `json:"data"`
	}

	PaystackEvent string
)

var (
	PaystackEventChargeSuccess  PaystackEvent = "charge.success"
	PaystackEventTransferFailed PaystackEvent = "transfer.failed"
)
