package mongo

import (
	"context"
	"iap/config"
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MG *MongoInstance

type mongodb struct {
	DbName   string
	MongoURL string
}

//ConnectToDB func connects to mongo db
func (MG *MongoInstance) ConnectToDB() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	mg := &mongodb{}
	mapstructure.Decode(cfg, &mg)
	client, err := mongo.NewClient(options.Client().ApplyURI(mg.MongoURL))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	db := client.Database(mg.DbName)
	MG = &MongoInstance{
		Client: client,
		DB:     db,
	}
	return nil

}
