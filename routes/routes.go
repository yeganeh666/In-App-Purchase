package routes

import (
	"iap/handlers"
	"iap/services/google"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	app := &google.App{
		DefaultHTTPClient:       http.DefaultClient,
		PubsubVerificationToken: os.Getenv("PUBSUB_VERIFICATION_TOKEN"),
	}
	route := mux.NewRouter()
	route.HandleFunc("/iap/{provider}/verify", handlers.Verify).Methods("POST")
	route.HandleFunc("/iap/google/acknowledgeSubscription", handlers.AcknowledgeSubscription).Methods("POST")
	route.HandleFunc("/pubsub/message/list", app.ListHandler)
	route.HandleFunc("/pubsub/message/receive", app.ReceiveMessagesHandler)

	return route
}