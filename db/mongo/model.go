package mongo

import "go.mongodb.org/mongo-driver/mongo"

//MongoInstance is mongoDB model
type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}
