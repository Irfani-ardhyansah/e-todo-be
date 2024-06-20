package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"errors"
)

type TaskRepositoryImpl struct{}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task {
	SQL := "INSERT INTO task(name, status) values(?, ?)"
	result, err := tx.ExecContext(ctx, SQL, task.Name, task.Status)
	helper.PanifIfError(err)

	id, err := result.LastInsertId()
	helper.PanifIfError(err)

	task.Id = int(id)
	return task
}

func (repository *TaskRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task {
	SQL := "UPDATE task SET name = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, task.Name, task.Status, task.Id)
	helper.PanifIfError(err)

	return task
}

func (repository *TaskRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, task domain.Task) {
	SQL := "DELETE FROM task WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, task.Id)
	helper.PanifIfError(err)
}

func (repository *TaskRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, taskId int) (domain.Task, error) {
	SQL := "SELECT id, name, status, created_at WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, taskId)
	helper.PanifIfError(err)
	defer rows.Close()

	task := domain.Task{}
	if rows.Next() {
		err := rows.Scan(&task.Id, &task.Name, &task.Status, &task.CreatedAt)
		helper.PanifIfError(err)
		return task, nil
	} else {
		return task, errors.New("Task Is Not Found")
	}
}

func (repository *TaskRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Task {
	SQL := "SELECT id, name, status, created_at"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanifIfError(err)
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		task := domain.Task{}
		err := rows.Scan(&task.Id, &task.Name, &task.Status, &task.CreatedAt)
		helper.PanifIfError(err)
		tasks = append(tasks, task)
	}

	return tasks
}
