package service

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"e-todo/model/web"
	"e-todo/repository"

	"github.com/go-playground/validator/v10"
)

type TimerServiceImpl struct {
	TimerRepository repository.TimerRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewTimerService(timerRepository repository.TimerRepository, DB *sql.DB, validate *validator.Validate) TimerService {
	return &TimerServiceImpl{
		TimerRepository: timerRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *TimerServiceImpl) Start(ctx context.Context, request web.TimerCreateRequest) web.TimerResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	task := domain.Timer{
		TaskId: request.TaskId,
		Time:   request.Time,
		Status: request.Status,
	}

	task = service.TimerRepository.Save(ctx, tx, task)

	return helper.ToTimerResponse(task)
}
