package models

type ReceiptData struct {
	ReceiptData string `json:"receiptData" validate:"required"`
}
