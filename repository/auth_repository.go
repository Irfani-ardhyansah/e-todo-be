package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type AuthRepository interface {
	SaveToken(ctx context.Context, tx *sql.Tx, userToken domain.UserToken) error
	CheckValidToken(ctx context.Context, db *sql.DB, userId int, refreshToken string) bool
}
