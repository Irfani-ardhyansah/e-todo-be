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
	TimerRepository         repository.TimerRepository
	DB                      *sql.DB
	Validate                *validator.Validate
	TimerHistoryReposeitory repository.TimerHistoryRepository
}

func NewTimerService(timerRepository repository.TimerRepository, DB *sql.DB, validate *validator.Validate, timerHistoryRepository repository.TimerHistoryRepository) TimerService {
	return &TimerServiceImpl{
		TimerRepository:         timerRepository,
		DB:                      DB,
		Validate:                validate,
		TimerHistoryReposeitory: timerHistoryRepository,
	}
}

func (service *TimerServiceImpl) Start(ctx context.Context, request web.TimerCreateRequest) web.TimerResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	timer := domain.Timer{
		TaskId: request.TaskId,
		Time:   request.Time,
		Status: request.Status,
	}

	timer = service.TimerRepository.Save(ctx, tx, timer)

	timerHistory := domain.TimerHistory{
		TimerId: timer.Id,
		TimeLog: request.Time,
		Status:  request.Status,
	}

	service.TimerHistoryReposeitory.Save(ctx, tx, timerHistory)

	return helper.ToTimerResponse(timer)
}

func (service *TimerServiceImpl) Update(ctx context.Context, request web.TimerUpdateRequest) web.TimerResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	timer := domain.Timer{
		Id:     request.Id,
		Time:   request.Time,
		Status: request.Status,
	}

	timer = service.TimerRepository.Update(ctx, tx, timer)

	timerHistory := domain.TimerHistory{
		TimerId: request.Id,
		TimeLog: request.Time,
		Status:  request.Status,
	}

	service.TimerHistoryReposeitory.Save(ctx, tx, timerHistory)

	return helper.ToTimerResponse(timer)
}
