package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"time"
)

type TimerHistoryRepositoryImpl struct{}

var currentTimeTimerHistory = time.Now().Format("2006-01-02 15:04:05")

func NewTimerHistoryRepository() TimerHistoryRepository {
	return &TimerHistoryRepositoryImpl{}
}

func (repository *TimerHistoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, timer domain.TimerHistory) {
	SQL := "INSERT INTO timer_histories(timer_id, time_log, status,  created_at) values(?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, timer.TimerId, timer.TimeLog, timer.Status, currentTimeTimerHistory)
	helper.PanifIfError(err)
}
