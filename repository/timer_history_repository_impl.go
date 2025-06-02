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

func NewTimerHistoryRepository() TimerHistoryRepository {
	return &TimerHistoryRepositoryImpl{}
}

func (repository *TimerHistoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, timer domain.TimerHistory) {
	SQL := "INSERT INTO timer_histories(timer_id, time_log, status,  created_at) values(?, ?, ?, ?)"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := tx.ExecContext(ctx, SQL, timer.TimerId, timer.TimeLog, timer.Status, currentTime)
	helper.PanifIfError(err)
}

func (repository *TimerHistoryRepositoryImpl) FindByParentId(ctx context.Context, db *sql.DB, timerId int) (domain.RelationTimerHistories, error) {
	SQL := "SELECT id, task_id, timer, status FROM timers WHERE id = ?"
	rows, err := db.QueryContext(ctx, SQL, timerId)
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
		return timer, errors.New("Task Is Not Found")
	}
}

func (repository *TimerHistoryRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) ([]domain.TaskDetail, error) {
	SQL := "SELECT tasks.id, name, timer, DATE(timers.created_at) as date FROM tasks JOIN timers ON tasks.id = timers.task_id ORDER BY timers.created_at DESC"
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanifIfError(err)
	defer rows.Close()

	tasks := []domain.TaskDetail{}
	for rows.Next() {
		task := domain.TaskDetail{}
		var timeString, dateString string
		err := rows.Scan(&task.Id, &task.TaskName, &timeString, &dateString)
		task.Time = timeString
		helper.PanifIfError(err)
		task.Date, err = time.Parse("2006-01-02", dateString)
		helper.PanifIfError(err)

		tasks = append(tasks, task)
	}

	return tasks, err
}
