package main

import (
	"context"
	"iap/db"
	"iap/db/mongo"
	"iap/routes"
	"iap/validators"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

var (
	router *mux.Router
)

func init() {
	validators.Init()
	router = routes.InitRoutes()
	var DB = map[string]db.DataBase{
		"mongo": &mongo.MongoInstance{},
	}
	DB["mongo"].ConnectToDB()
}

func main() {
	l := hclog.Default()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		l.Info("system call:", oscall)
		cancel()
	}()

	if err := routes.Serve(ctx, router, l); err != nil {
		l.Error("failed to serve:", "error", err)
	}
}
