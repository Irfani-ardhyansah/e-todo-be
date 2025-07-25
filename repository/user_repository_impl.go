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
	SQL := "SELECT id, email, name, password FROM users WHERE email = ?"
	rows, err := db.Query(SQL, email)
	helper.PanifIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		helper.PanifIfError(err)
		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, db *sql.DB, userId int) (domain.User, error) {
	SQL := "SELECT id, email, name, password FROM users WHERE id = ?"
	rows, err := db.Query(SQL, userId)
	helper.PanifIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		helper.PanifIfError(err)
		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}

func (repository *UserRepositoryImpl) SaveToken(ctx context.Context, tx *sql.Tx, userToken domain.UserToken) error {
	SQL := "INSERT INTO user_tokens(user_id, refresh_token, is_valid) values(?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, userToken.UserId, userToken.RefreshToken, userToken.IsValid)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (repository *UserRepositoryImpl) CheckValidToken(ctx context.Context, db *sql.DB, userId int, refreshToken string) bool {
	SQL := "SELECT is_valid FROM user_tokens WHERE user_id = ? AND refresh_token = ?"
	rows, err := db.Query(SQL, userId, refreshToken)

	helper.PanifIfError(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}
