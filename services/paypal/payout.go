package paypal

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/plutov/paypal"
)

func (p *PayPal) Pay(amount string) (*PayPalResponse, error) {
	p.Client, _ = paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
	accessToken, err := p.Client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	request := paypalreq(amount)
	buf := bytes.NewBuffer(EncodeToBytes(request))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", p.Client.APIBase, "/v1/payments/payment"), buf)
	if err != nil {
		return nil, err
	}
	header := fmt.Sprintf("Bearer %v", accessToken)
	req.Header.Set("Content-Type", " application/json")
	req.Header.Set("Authorization", header)
	res := &paypal.PayoutResponse{}
	err = p.Client.SendWithBasicAuth(req, res)
	return &PayPalResponse{
		TransactionID: res.Items[0].TransactionID,
		Href:          res.Links[0].Href,
	}, nil
}
func paypalreq(amount string) *PayPalRequest {
	file, _ := ioutil.ReadFile("test.json")
	data := &PayPalRequest{}
	_ = json.Unmarshal([]byte(file), &data)
	data.Transactions[0].Amount.Total = amount
	return data
	// return &PayPalRequest{
	// 	Intent: "sale",
	// 	Payer: {
	// 		PaymentMethod: "paypal",
	// 	},
	// 	[]Transactions{
	// 		Amount{
	// 			Total:    "436",
	// 			Currency: "USD",
	// 		},
	// 		Description:   "This is the payment transaction description.",
	// 		Custom:        "EBAY_EMS_90048630024435",
	// 		InvoiceNumber: "48787589673",
	// 		PaymentOptions{
	// 			AllowedPaymentMethod: "INSTANT_FUNDING_SOURCE",
	// 		},
	// 	},
	// 	NoteToPayer: "Contact us for any questions on your order.",
	// 	RedirectUrls{
	// 		ReturnURL: "https://example.com",
	// 		CancelURL: "https://example.com",
	// 	},
	// }

}
func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(p)

	return buf.Bytes()
}
