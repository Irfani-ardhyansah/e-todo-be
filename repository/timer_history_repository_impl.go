package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"errors"
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

func (repository *TimerHistoryRepositoryImpl) FindByParentId(ctx context.Context, db *sql.DB, timerId int) (domain.RelationTimerHistories, error) {
	SQL := "SELECT id, task_id, timer, status FROM timers WHERE id = ?"
	rows, err := db.Query(SQL, timerId)
	helper.PanifIfError(err)
	defer rows.Close()

	timer := domain.RelationTimerHistories{}
	if rows.Next() {
		err := rows.Scan(&timer.Id, &timer.TaskId, &timer.Time, &timer.Status)
		helper.PanifIfError(err)
		timeHistory := domain.TimerHistory{}
		SQL1 := "SELECT id, status, time_log, created_at FROM timer_histories WHERE timer_id = ?"
		rows1, err := db.Query(SQL1, timerId)
		helper.PanifIfError(err)
		defer rows1.Close()

		for rows1.Next() {
			err := rows1.Scan(&timeHistory.Id, &timeHistory.Status, &timeHistory.TimeLog, &timeHistory.CreatedAt)
			helper.PanifIfError(err)

			timer.Histories = append(timer.Histories, timeHistory)
		}

		return timer, nil
	} else {
		return domain.RelationTimerHistories{}, errors.New("Task Is Not Found")
	}
}
