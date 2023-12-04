// main.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MidtransPaymentRequest struct {
	VaNumbers []struct {
		VaNumber string `json:"va_number"`
		Bank     string `json:"bank"`
	} `json:"va_numbers"`
	TransactionTime   string        `json:"transaction_time"`
	TransactionStatus string        `json:"transaction_status"`
	TransactionID     string        `json:"transaction_id"`
	StatusMessage     string        `json:"status_message"`
	StatusCode        string        `json:"status_code"`
	SignatureKey      string        `json:"signature_key"`
	SettlementTime    string        `json:"settlement_time"`
	PaymentType       string        `json:"payment_type"`
	PaymentAmounts    []interface{} `json:"payment_amounts"`
	OrderID           string        `json:"order_id"`
	MerchantID        string        `json:"merchant_id"`
	GrossAmount       string        `json:"gross_amount"`
	FraudStatus       string        `json:"fraud_status"`
	ExpiryTime        string        `json:"expiry_time"`
	Currency          string        `json:"currency"`
}

func main() {
	e := echo.New()

	// Define the webhook route
	e.POST("/webhook", handleWebhook)

	// Start the server
	err := e.Start(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handleWebhook(c echo.Context) error {
	fmt.Println("Received webhook payload")

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Process the webhook payload here
	request := &MidtransPaymentRequest{}
	err = json.Unmarshal(body, request)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	fmt.Println("UPDATE transactions SET transaction_status = ", request.TransactionStatus, " WHERE id = ", request.OrderID)

	// Respond to the webhook request
	return c.String(http.StatusOK, "Webhook received successfully")
}
