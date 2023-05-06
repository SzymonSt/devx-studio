package util

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(dbUrl string) (dbClient *mongo.Client) {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	err = dbClient.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	return
}
