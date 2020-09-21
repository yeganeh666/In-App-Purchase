package paypal

import "github.com/plutov/paypal"

type PayPal struct {
	Client *paypal.Client
}
type PayPalResponse struct {
	TransactionID string `json:"transaction_id"`
	Href          string `json:"href"`
}
type PayPalRequest struct {
	Intent string `json:"intent"`
	Payer  struct {
		PaymentMethod string `json:"payment_method"`
	} `json:"payer"`
	Transactions []struct {
		Amount struct {
			Total    string `json:"total"`
			Currency string `json:"currency"`
			Details  struct {
				Subtotal         string `json:"subtotal"`
				Tax              string `json:"tax"`
				Shipping         string `json:"shipping"`
				HandlingFee      string `json:"handling_fee"`
				ShippingDiscount string `json:"shipping_discount"`
				Insurance        string `json:"insurance"`
			} `json:"details"`
		} `json:"amount"`
		Description    string `json:"description"`
		Custom         string `json:"custom"`
		InvoiceNumber  string `json:"invoice_number"`
		PaymentOptions struct {
			AllowedPaymentMethod string `json:"allowed_payment_method"`
		} `json:"payment_options"`
		SoftDescriptor string `json:"soft_descriptor"`
		ItemList       struct {
			Items []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Quantity    string `json:"quantity"`
				Price       string `json:"price"`
				Tax         string `json:"tax"`
				Sku         string `json:"sku"`
				Currency    string `json:"currency"`
			} `json:"items"`
			ShippingAddress struct {
				RecipientName string `json:"recipient_name"`
				Line1         string `json:"line1"`
				Line2         string `json:"line2"`
				City          string `json:"city"`
				CountryCode   string `json:"country_code"`
				PostalCode    string `json:"postal_code"`
				Phone         string `json:"phone"`
				State         string `json:"state"`
			} `json:"shipping_address"`
		} `json:"item_list"`
	} `json:"transactions"`
	NoteToPayer  string `json:"note_to_payer"`
	RedirectUrls struct {
		ReturnURL string `json:"return_url"`
		CancelURL string `json:"cancel_url"`
	} `json:"redirect_urls"`
}
