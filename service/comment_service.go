package service

import (
	"context"
	"e-todo/model/web"
)

type CommentService interface {
	Create(ctx context.Context, task web.CommentCreateRequest) web.CommentResponse
	Update(ctx context.Context, task web.CommentUpdateRequest) web.CommentResponse
	Delete(ctx context.Context, commentId int)
	FindById(ctx context.Context, commentId int) web.CommentResponse
	FindAll(ctx context.Context) []web.CommentResponse
}
