package service

import (
	"context"
	"database/sql"
	exception "e-todo/excception"
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/repository"

	"github.com/go-playground/validator/v10"
)

type TimerHistoryServiceImpl struct {
	TimerHistoryRepository repository.TimerHistoryRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewTimerHistoryService(timerHistoryRepository repository.TimerHistoryRepository, DB *sql.DB, validate *validator.Validate) TimerHistoryService {
	return &TimerHistoryServiceImpl{
		TimerHistoryRepository: timerHistoryRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *TimerHistoryServiceImpl) FindByParentId(ctx context.Context, timerId int) web.RelationTimerHistoriesResponse {
	// tx, err := service.DB.Begin()
	// helper.PanifIfError(err)
	// defer helper.CommitOrRollback(tx)

	timer, err := service.TimerHistoryRepository.FindByParentId(ctx, service.DB, timerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRelationTimerHistoriesResponse(timer)
}
