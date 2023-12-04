package main

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

const (
	midtransBaseURL  = "https://api.sandbox.midtrans.com/snap/v1"
	clientKey        = "yourClientKey"
	serverKey        = "yourServerKey"
	orderID          = "WEEBO-999"
	transactionTotal = 50000
)

type MidtransPaymentRequest struct {
	TransactionDetails MidtransTransactionDetails `json:"transaction_details"`
}

type MidtransTransactionDetails struct {
	OrderID  string `json:"order_id"`
	GrossAmt int    `json:"gross_amount"`
}

func main() {
	snapClient := snap.Client{}

	snapClient.New(serverKey, midtrans.Sandbox)

	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: transactionTotal,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "Tukimin",
			LName: "Pergi Ke pasar",
			Email: "felixajax@weeboo.com",
		},
	}

	snapResponse, _ := snapClient.CreateTransaction(request)

	fmt.Println("Response:", snapResponse)
}

type Transactions struct {
	ID                int64
	TransactionTitle  string
	TransactionAmount int
	TransactionStatus string
}
