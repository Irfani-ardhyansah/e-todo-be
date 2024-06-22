package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type TimerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, timer domain.Timer) domain.Timer
}
