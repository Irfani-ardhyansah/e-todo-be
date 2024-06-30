package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
