package service

import (
	"context"
	"golang-database-user/model"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.MstUser) model.MstUser
	UpdateUser(ctx context.Context, user model.MstUser, userId string) model.MstUser
	ReadUser(ctx context.Context) ([]model.MstUser, error)
	DeleteUser(ctx context.Context, userId string) error
}
