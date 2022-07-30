package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnection() *mongo.Database {
	// Todo: add to env var
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"))
	if err != nil {
		log.Fatal(err)
	}
	// Todo- review if WithTimeout is essential
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer func(client *mongo.Client, ctx context.Context) {
	//	err := client.Disconnect(ctx)
	//	if err != nil {
	//		log.Print("client disconnection defer failed")
	//	}
	//}(client, ctx)
	return client.Database("Faas")
}
