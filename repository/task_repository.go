package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type TaskRepository interface {
	Save(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task
	Update(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task
	Delete(ctx context.Context, tx *sql.Tx, task domain.Task)
	FindById(ctx context.Context, tx *sql.Tx, taskId int) (domain.Task, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Task
}
