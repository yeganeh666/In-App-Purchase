package mongo

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func (MG *MongoInstance) GetAmount(ItemID []string) string {
	var Items []interface{}
	for _, id := range ItemID {
		docs := []interface{}{
			bson.D{
				{"ID", id},
			}}
		Items = append(Items, docs)
	}
	type item struct {
		id    int    `bson:"ID"`
		price string `bson:"price"`
	}
	var result []item
	cursor, err := MG.DB.Collection("items").Find(context.TODO(), Items)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		r := item{}
		if err = cursor.Decode(&r); err != nil {
			log.Fatal(err)
		}
		result = append(result, r)
	}
	var sum float64
	for _, i := range result {
		price, _ := strconv.Atoi(i.price)
		sum += float64(price)
	}
	return fmt.Sprintf("%f", sum)
}
func (MG *MongoInstance) InsertTransactions(TransactionID string) error {

	_, err := MG.DB.Collection("transactions").InsertOne(context.TODO(), bson.D{{"transactionID", TransactionID}})
	if err != nil {
		return err
	}
	return nil
}
