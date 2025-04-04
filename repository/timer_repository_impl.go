package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"time"
)

type TimerRepositoryImpl struct{}

var currentTimeTimer = time.Now().Format("2006-01-02 15:04:05")

func NewTimerRepository() TimerRepository {
	return &TimerRepositoryImpl{}
}

func (repository *TimerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, timer domain.Timer) domain.Timer {
	SQL := "INSERT INTO timers(task_id, timer, title, status, created_at) values(?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, timer.TaskId, timer.Time, timer.Title, timer.Status, currentTimeTask)
	helper.PanifIfError(err)

	id, err := result.LastInsertId()
	helper.PanifIfError(err)

	timer.Id = int(id)
	timer.CreatedAt = currentTimeTask
	return timer
}

func (repository *TimerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, timer domain.Timer) domain.Timer {
	SQL := "UPDATE timers SET timer = ?, title = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, timer.Time, timer.Title, timer.Status, timer.Id)
	helper.PanifIfError(err)

	return timer
}
