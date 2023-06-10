package mongoRepository

import (
	"context"
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAdapter struct {
	conn *mongo.Database
}

func NewMongoAdapter(conn *mongo.Database) *MongoAdapter {
	return &MongoAdapter{
		conn: conn,
	}
}

func (adapter *MongoAdapter) Save(user domain.User) error {
	result, err := adapter.conn.Collection("users").InsertOne(context.TODO(), user)

	if err != nil {
		return fmt.Errorf("error saving new user user %s", err)
	}

	fmt.Println(result)

	return nil
}

func (adapter *MongoAdapter) FindByEmail(email string) (domain.User, error) {
	var result domain.User

	err := adapter.conn.Collection("users").FindOne(context.TODO(), bson.M{
		"email": email,
	}).Decode(&result)

	if err != nil {
		return domain.User{}, fmt.Errorf("error searching user by email: %s", err.Error())
	}

	return result, nil
}
