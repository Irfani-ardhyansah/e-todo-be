package service

import (
	"context"
	"e-todo/model/web"
)

type UserService interface {
	Create(ctx context.Context, task web.UserCreateRequest) web.UserResponse
}
