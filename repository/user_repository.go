package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.User, error)
	FindById(ctx context.Context, db *sql.DB, userId int) (domain.User, error)
	SaveToken(ctx context.Context, tx *sql.Tx, userToken domain.UserToken) error
	CheckValidToken(ctx context.Context, db *sql.DB, userId int, refreshToken string) bool
}
