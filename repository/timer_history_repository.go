package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type TimerHistoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, timerHistory domain.TimerHistory)
}
