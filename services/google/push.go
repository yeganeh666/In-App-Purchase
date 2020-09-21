package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"google.golang.org/api/pubsub/v1"
)

type PubSub struct {
	l hclog.Logger
	// pubsubVerificationToken is a shared secret between the the publisher of
	// the message and this application.
	PubSubVerificationToken string

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

// NewPubSubNotificationService returns new pubsub instance.
func NewPubSubNotificationService(l hclog.Logger, pubSubVerificationToken string) *PubSub {
	return &PubSub{
		l:                       l,
		DefaultHTTPClient:       http.DefaultClient,
		PubSubVerificationToken: pubSubVerificationToken,
	}
}

// MessageHandler receive incomming messages from google pub/sub service
// push mode
func (ps *PubSub) MessageHandler(w http.ResponseWriter, r *http.Request) {
	ps.l.Info("message received from google ...")

	var pr pushRequest
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	ps.MessagesMu.Lock()
	defer ps.MessagesMu.Unlock()
	// Limit to ten.
	ps.Messages = append(ps.Messages, pr.Message.Data)
	if len(ps.Messages) > maxMessages {
		ps.Messages = ps.Messages[len(ps.Messages)-maxMessages:]
	}

	fmt.Fprint(w, "OK")
}

// VerifyRequest verifies request that come from google and check credentials.
func (ps *PubSub) VerifyRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that the request originates from the application.
		// a.pubsubVerificationToken = os.Getenv("PUBSUB_VERIFICATION_TOKEN")
		if token, ok := r.URL.Query()["token"]; !ok || len(token) != 1 || token[0] != ps.PubSubVerificationToken {
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
		v, err := idtoken.NewValidator(r.Context(), option.WithHTTPClient(ps.DefaultHTTPClient))
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
		next.ServeHTTP(w, r)
	})
}

func (ps *PubSub) ListHandler(w http.ResponseWriter, r *http.Request) {
	ps.MessagesMu.Lock()
	defer ps.MessagesMu.Unlock()

	fmt.Fprintln(w, "Messages:")
	for _, v := range ps.Messages {
		fmt.Fprintf(w, "Message: %v\n", v)
	}
}
