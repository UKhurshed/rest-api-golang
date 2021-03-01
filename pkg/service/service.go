package service

import (
	"context"
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user todo.User) error
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(token string) (string, error)
}
 type Service struct {
 	Authorization
 }

 func NewService (repos *repository.Repository) *Service{
 	return &Service{
 		Authorization: NewAuthService(repos.Authorization),
	}
 }