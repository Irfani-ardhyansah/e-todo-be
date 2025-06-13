package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type TimerHistoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, timerHistory domain.TimerHistory)
	FindByParentId(ctx context.Context, tx *sql.DB, timerId int) (domain.RelationTimerHistories, error)
	GetAll(ctx context.Context, tx *sql.DB) ([]domain.TaskDetail, error)
	GetByTaskId(ctx context.Context, tx *sql.DB, taskId int) ([]domain.TaskDetail, error)
}
