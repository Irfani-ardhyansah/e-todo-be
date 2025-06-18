package repository

import (
	"context"
	"database/sql"
	"e-todo/model/domain"
)

type CommentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Comment domain.Comment) domain.Comment
	Update(ctx context.Context, tx *sql.Tx, Comment domain.Comment) domain.Comment
	Delete(ctx context.Context, tx *sql.Tx, Comment domain.Comment)
	FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.Comment, error)
	FindAll(ctx context.Context, tx *sql.Tx, taskId int) []domain.Comment
}
