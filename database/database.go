package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase(user, password, host, port, dbname string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", user, password, host, port, dbname, user)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
