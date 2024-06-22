package service

import (
	"context"
	"e-todo/model/web"
)

type TimerService interface {
	Start(ctx context.Context, task web.TimerCreateRequest) web.TimerResponse
}
