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
		return fmt.Errorf("error saving new user %s", err)
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

	if err == mongo.ErrNoDocuments {
		return domain.User{}, mongo.ErrNoDocuments
	}

	return result, nil
}

func (adapter *MongoAdapter) UpdatePassword(email, newPassword string) error {
	filters := bson.M{"email": email}
	update := bson.M{"password": newPassword}

	updateResult, err := adapter.conn.Collection("users").UpdateOne(context.TODO(), filters, update)

	if err != nil {
		return err
	}

	fmt.Println(updateResult)

	return nil
}

func (adapter *MongoAdapter) VerifyResetPassword(email string) (domain.PasswordReset, error) {
	var requestPasswordReset domain.PasswordReset

	filters := bson.M{"active": true, "email": email}
	err := adapter.conn.Collection("users_change_password").FindOne(context.TODO(), filters).Decode(&requestPasswordReset)

	if err != nil {
		return requestPasswordReset, err
	}

	if err == mongo.ErrNoDocuments {
		return requestPasswordReset, mongo.ErrNoDocuments
	}

	return requestPasswordReset, nil
}

func (adapter *MongoAdapter) InactiveOtp(email string) error {
	filters := bson.M{"active": true, "email": email}
	update := bson.M{"active": false}
	result, err := adapter.conn.Collection("users_change_password").UpdateOne(context.TODO(), filters, update)

	if err != nil {
		return fmt.Errorf("error triying update document")
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not matched document")
	}

	return nil
}

func (adapter *MongoAdapter) GenerateOtp(data domain.PasswordReset) error {
	result, err := adapter.conn.Collection("users_change_password").InsertOne(context.TODO(), bson.M{
		"code":   data.Code,
		"active": data.Active,
		"email":  data.Email,
	})

	if err != nil {
		return fmt.Errorf("error saving new otp entry %s", err)
	}

	fmt.Println(result)

	return nil
}
