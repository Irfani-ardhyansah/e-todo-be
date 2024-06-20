package service

import (
	"context"
	"e-todo/model/web"
)

type TaskService interface {
	Create(ctx context.Context, task web.TaskCreateRequest) web.TaskResponse
	Update(ctx context.Context, task web.TaskUpdateRequest) web.TaskResponse
	Delete(ctx context.Context, taskId int)
	FindById(ctx context.Context, taskId int) web.TaskResponse
	FindAll(ctx context.Context) []web.TaskResponse
}
