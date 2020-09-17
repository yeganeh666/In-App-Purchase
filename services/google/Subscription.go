package google

import (
	"context"
	"io/ioutil"

	"github.com/awa/go-iap/playstore"
	"github.com/fatih/structs"
	"google.golang.org/api/androidpublisher/v3"
)

//Google IAB
type Google struct {
	Client           *playstore.Client
	Package          string `json:"package" validate:"required"`
	SubscriptionID   string `json:"subscriptionID" validate:"required"`
	PurchaseToken    string `json:"purchaseToken" validate:"required"`
	DeveloperPayload string `json:"developerPayload"`
}

//Verify Subscription in Google
func (g *Google) Verify() map[string]interface{} {
	// You need to prepare a public key for your Android app's in app billing
	// at https://console.developers.google.com.
	jsonKey, err := ioutil.ReadFile("json_Key.json")
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
	resp, err := g.Client.VerifySubscription(ctx, g.Package, g.SubscriptionID, g.PurchaseToken)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error()}
	}
	m := structs.Map(resp)
	return m
}

//AcknowledgeSubscription acknowledges a subscription purchase.
func (g *Google) AcknowledgeSubscription() error {
	jsonKey, err := ioutil.ReadFile("json_Key.json")
	if err != nil {
		return err
	}
	client, err := playstore.New(jsonKey)
	if err != nil {
		return err
	}
	g.Client = client
	spa := &androidpublisher.SubscriptionPurchasesAcknowledgeRequest{
		DeveloperPayload: g.DeveloperPayload,
	}
	ctx := context.Background()
	err = g.Client.AcknowledgeSubscription(ctx, g.Package, g.SubscriptionID, g.PurchaseToken, spa)
	if err != nil {
		return err
	}
	return err
}
