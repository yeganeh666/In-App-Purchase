package models

type VerifySubscription struct {
	Package        string `json:"package" validate:"required"`
	SubscriptionID string `json:"subscriptionID" validate:"required"`
	PurchaseToken  string `json:"purchaseToken" validate:"required"`
}
