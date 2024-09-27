package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
)

type AuthRepositoryImpl struct{}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) SaveToken(ctx context.Context, tx *sql.Tx, userToken domain.UserToken) error {
	SQLDelete := "DELETE FROM user_tokens WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, SQLDelete, userToken.UserId)
	if err != nil {
		return err
	}

	SQL := "INSERT INTO user_tokens(user_id, refresh_token, is_valid) values(?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, userToken.UserId, userToken.RefreshToken, userToken.IsValid)
	if err != nil {
		return err
	}

	return nil
}

func (repository *AuthRepositoryImpl) CheckValidToken(ctx context.Context, db *sql.DB, userId int, refreshToken string) bool {
	SQL := "SELECT is_valid FROM user_tokens WHERE user_id = ? AND refresh_token = ?"
	rows, err := db.Query(SQL, userId, refreshToken)
	helper.PanifIfError(err)
	defer rows.Close()

	if rows.Next() {
		return true
	}

	return false
}
