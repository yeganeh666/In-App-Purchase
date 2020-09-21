package routes

import (
	"iap/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(token string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/iap/{provider}/verify", handlers.Verify).Methods("POST")
	router.HandleFunc("/iap/google/acknowledgeSubscription", handlers.AcknowledgeSubscription).Methods("POST")
	//router.HandleFunc("/pubsub/message/list", app.ListHandler)
	//router.HandleFunc("/pubsub/message/receive", app.ReceiveMessagesHandler)

	return router
}
