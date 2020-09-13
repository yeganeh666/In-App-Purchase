package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/awa/go-iap/appstore"
	"github.com/awa/go-iap/playstore"
	"github.com/fatih/structs"
)

//Services for handle IAPs
type Services interface {
	NewService()
	Verify(map[string]interface{}) map[string]interface{}
}

var X Services

//Apple IAP
type Apple struct {
	Client *appstore.Client
}

//NewService create new Apple client
func (a *Apple) NewService() {
	a.Client = appstore.New()
	fmt.Println("OK")
	return
}

//Verify ReceiptData in Apple
func (a *Apple) Verify(request map[string]interface{}) map[string]interface{} {
	req := appstore.IAPRequest{
		ReceiptData: request["ReceiptData"].(string),
	}
	resp := &appstore.IAPResponse{}
	ctx := context.Background()
	err := a.Client.Verify(ctx, req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
	m := structs.Map(resp)
	return m
}

//Google IAB
type Google struct {
	Client *playstore.Client
}

//NewService create new Google client
func (g *Google) NewService() {
	// You need to prepare a public key for your Android app's in app billing
	// at https://console.developers.google.com.
	jsonKey, err := ioutil.ReadFile("jsonKey.json")
	if err != nil {
		log.Fatal(err)
	}
	client, err := playstore.New(jsonKey)
	if err != nil {
		log.Fatal(err)
	}
	g.Client = client
	fmt.Println("OK")
	return
}

//Verify Subscription in Google
func (g *Google) Verify(request map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	resp, err := g.Client.VerifySubscription(ctx, request["package"].(string), request["subscriptionID"].(string), request["purchaseToken"].(string))
	if err != nil {
		log.Fatal(err)
	}
	m := structs.Map(resp)
	fmt.Println("OK")
	return m
}
