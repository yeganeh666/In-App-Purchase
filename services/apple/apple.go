package apple

import (
	"context"

	"github.com/awa/go-iap/appstore"
	"github.com/fatih/structs"
)

//Apple IAP
type Apple struct {
	Client      *appstore.Client
	ReceiptData string `json:"receiptData" validate:"required"`
}

//Verify ReceiptData in Apple
func (a *Apple) Verify() map[string]interface{} {
	a.Client = appstore.New()
	req := appstore.IAPRequest{
		ReceiptData: a.ReceiptData,
	}
	resp := &appstore.IAPResponse{}
	ctx := context.Background()
	err := a.Client.Verify(ctx, req, resp)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error()}
	}
	m := structs.Map(resp)
	return m
}
