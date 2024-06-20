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

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTaskService(taskRepository repository.TaskRepository, DB *sql.DB, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TaskServiceImpl) Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	task := domain.Task{
		Name:   request.Name,
		Status: request.Status,
	}

	task = service.TaskRepository.Save(ctx, tx, task)

	return helper.ToTaskResponse(task)
}

func (service *TaskServiceImpl) Update(ctx context.Context, request web.TaskUpdateRequest) web.TaskResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	task.Name = request.Name
	task.Status = request.Status
	task = service.TaskRepository.Update(ctx, tx, task)
	return helper.ToTaskResponse(task)
}

func (service *TaskServiceImpl) Delete(ctx context.Context, taskId int) {
	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindById(ctx, tx, taskId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TaskRepository.Delete(ctx, tx, task)
}

func (service *TaskServiceImpl) FindById(ctx context.Context, taskId int) web.TaskResponse {
	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindById(ctx, tx, taskId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTaskResponse(task)
}

func (service *TaskServiceImpl) FindAll(ctx context.Context) []web.TaskResponse {
	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)
	tasks := service.TaskRepository.FindAll(ctx, tx)

	return helper.ToTaskResponses(tasks)
}
