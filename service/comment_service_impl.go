package service

import (
	"context"
	"database/sql"
	exception "e-todo/excception"
	"e-todo/helper"
	"e-todo/model/domain"
	"e-todo/model/web"
	"e-todo/repository"

	"github.com/go-playground/validator/v10"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCommentService(commentRepository repository.CommentRepository, DB *sql.DB, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CommentServiceImpl) Create(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	task := domain.Comment{
		TaskId:   request.TaskId,
		UserId:   request.UserId,
		ParentId: request.ParentId,
		Comment:  request.Comment,
	}

	task = service.CommentRepository.Save(ctx, tx, task)

	return helper.ToCommentResponse(task)
}

func (service *CommentServiceImpl) Update(ctx context.Context, request web.CommentUpdateRequest) web.CommentResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	comment.Comment = request.Comment
	comment = service.CommentRepository.Update(ctx, tx, comment)
	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) Delete(ctx context.Context, commentId int) {
	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CommentRepository.Delete(ctx, tx, comment)
}

func (service *CommentServiceImpl) FindById(ctx context.Context, commentId int) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) FindAll(ctx context.Context) []web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)
	comments := service.CommentRepository.FindAll(ctx, tx)

	return helper.ToCommentResponses(comments)
}
