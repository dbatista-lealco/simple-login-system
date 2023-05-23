package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ConnectionContext context.Context = context.Background()
var Conn *mongo.Client

func NewConnection() (*mongo.Client, error) {
	if Conn != nil {
		return Conn, nil
	}
	// TODO: replace with OS environments
	mongoUri := "mongodb://localhost:27017"

	Conn, err := mongo.Connect(ConnectionContext, options.Client().ApplyURI(mongoUri))

	if err != nil {
		return nil, fmt.Errorf("error connection to mongo database: %s", err.Error())
	}

	return Conn, nil
}
