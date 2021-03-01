package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"todo"
)

type AuthMongo struct {
	db *mongo.Collection
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db.Collection("user")}
}

func (r *AuthMongo) CreateUser(ctx context.Context, user todo.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err

}

func (r *AuthMongo) GetUser(ctx context.Context, email, password string) (todo.User, error) {
	var user todo.User
	if err := r.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return todo.User{}, err
		}
		return todo.User{}, err
	}
	return user, nil
}
