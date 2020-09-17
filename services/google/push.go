package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"google.golang.org/api/pubsub/v1"
)

type App struct {
	// pubsubVerificationToken is a shared secret between the the publisher of
	// the message and this application.
	PubsubVerificationToken string

	// Messages received by this instance.
	MessagesMu sync.Mutex
	Messages   []string

	// defaultHTTPClient aliases http.DefaultClient for testing
	DefaultHTTPClient *http.Client
}
type pushRequest struct {
	Message      pubsub.PubsubMessage `json:"message"`
	Subscription string               `json:"subscription"`
}

const maxMessages = 10

func (a *App) ReceiveMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Verify that the request originates from the application.
	// a.pubsubVerificationToken = os.Getenv("PUBSUB_VERIFICATION_TOKEN")
	if token, ok := r.URL.Query()["token"]; !ok || len(token) != 1 || token[0] != a.PubsubVerificationToken {
		http.Error(w, "Bad token", http.StatusBadRequest)
		return
	}

	// Get the Cloud Pub/Sub-generated JWT in the "Authorization" header.
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(strings.Split(authHeader, " ")) != 2 {
		http.Error(w, "Missing Authorization header", http.StatusBadRequest)
		return
	}
	token := strings.Split(authHeader, " ")[1]
	// Verify and decode the JWT.
	// If you don't need to control the HTTP client used you can use the
	// convenience method idtoken.Validate instead of creating a Validator.
	v, err := idtoken.NewValidator(r.Context(), option.WithHTTPClient(a.DefaultHTTPClient))
	if err != nil {
		http.Error(w, "Unable to create Validator", http.StatusBadRequest)
		return
	}
	// Please change http://example.com to match with the value you are
	// providing while creating the subscription.
	payload, err := v.Validate(r.Context(), token, "http://example.com")
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Token: %v", err), http.StatusBadRequest)
		return
	}
	if payload.Issuer != "accounts.google.com" && payload.Issuer != "https://accounts.google.com" {
		http.Error(w, "Wrong Issuer", http.StatusBadRequest)
	}

	var pr pushRequest
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	a.MessagesMu.Lock()
	defer a.MessagesMu.Unlock()
	// Limit to ten.
	a.Messages = append(a.Messages, pr.Message.Data)
	if len(a.Messages) > maxMessages {
		a.Messages = a.Messages[len(a.Messages)-maxMessages:]
	}

	fmt.Fprint(w, "OK")
}

func (a *App) ListHandler(w http.ResponseWriter, r *http.Request) {
	a.MessagesMu.Lock()
	defer a.MessagesMu.Unlock()

	fmt.Fprintln(w, "Messages:")
	for _, v := range a.Messages {
		fmt.Fprintf(w, "Message: %v\n", v)
	}
}
