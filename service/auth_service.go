package service

import (
	"context"
	"e-todo/model/web"
)

type AuthService interface {
	Login(ctx context.Context, user web.UserCreateRequest) web.UserLoginResponse
	RefreshToken(ctx context.Context, refresh web.RefreshTokenRequest) web.UserLoginResponse
}
