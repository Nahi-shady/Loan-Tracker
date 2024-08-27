package bootstrap

import (
	"context"
	"log"

	"loan-tracker/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()

	mongodbURI := env.DBURI

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
