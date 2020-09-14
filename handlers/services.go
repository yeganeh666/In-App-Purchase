package handlers

import (
	"context"
	"io/ioutil"

	"github.com/awa/go-iap/appstore"
	"github.com/awa/go-iap/playstore"
	"github.com/fatih/structs"
)

//Services for handle IAPs
type Services interface {
	Verify(map[string]interface{}) map[string]interface{}
}

var X Services

//Apple IAP
type Apple struct {
	Client *appstore.Client
}

//Verify ReceiptData in Apple
func (a *Apple) Verify(request map[string]interface{}) map[string]interface{} {
	a.Client = appstore.New()
	req := appstore.IAPRequest{
		ReceiptData: request["ReceiptData"].(string),
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

//Google IAB
type Google struct {
	Client *playstore.Client
}

//Verify Subscription in Google
func (g *Google) Verify(request map[string]interface{}) map[string]interface{} {
	// You need to prepare a public key for your Android app's in app billing
	// at https://console.developers.google.com.
	jsonKey, err := ioutil.ReadFile("jsonKey.json")
	if err != nil {
		return map[string]interface{}{
			"error": err.Error()}
	}
	client, err := playstore.New(jsonKey)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error()}
	}
	g.Client = client
	ctx := context.Background()
	resp, err := g.Client.VerifySubscription(ctx, request["Package"].(string), request["SubscriptionID"].(string), request["PurchaseToken"].(string))
	if err != nil {
		return map[string]interface{}{
			"error": err.Error()}
	}
	m := structs.Map(resp)
	return m
}
