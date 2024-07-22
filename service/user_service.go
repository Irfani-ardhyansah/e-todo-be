package service

import (
	"context"
	"e-todo/model/web"
)

type UserService interface {
	Create(ctx context.Context, task web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, user web.UserCreateRequest) web.UserLoginResponse
	RefreshToken(ctx context.Context, id int, refreshToken string) web.UserLoginResponse
}
