package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"todo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user todo.User) error
	GetUser(ctx context.Context, email, password string) (todo.User, error)
}
type Repository struct {
	Authorization
}

func NewRepository (db *mongo.Database) *Repository{
	return &Repository{
			Authorization: NewAuthMongo(db),
	}
}
