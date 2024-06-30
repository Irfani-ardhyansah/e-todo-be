package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"time"
)

type UserRepositoryImpl struct{}

var currentTimeUser = time.Now().Format("2006-01-02 15:04:05")

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(email, password, created_at) values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, currentTimeUser)
	helper.PanifIfError(err)

	id, err := result.LastInsertId()
	helper.PanifIfError(err)

	user.Id = int(id)
	user.CreatedAt = currentTimeTask
	return user
}
