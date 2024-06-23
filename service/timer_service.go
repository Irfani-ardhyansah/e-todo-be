package service

import (
	"context"
	"e-todo/model/web"
)

type TimerService interface {
	Start(ctx context.Context, timer web.TimerCreateRequest) web.TimerResponse
	Update(ctx context.Context, imer web.TimerUpdateRequest) web.TimerResponse
}
