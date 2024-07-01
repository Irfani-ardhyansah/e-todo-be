package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"errors"
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

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.User, error) {
	SQL := "SELECT id, email, password FROM users WHERE email = ?"
	rows, err := db.Query(SQL, email)
	helper.PanifIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		helper.PanifIfError(err)
		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}
