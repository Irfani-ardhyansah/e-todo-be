package repository

import (
	"context"
	"database/sql"
	"e-todo/helper"
	"e-todo/model/domain"
	"errors"
)

type CommentRepositoryImpl struct{}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (repository *CommentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	SQL := "INSERT INTO comments(task_id, user_id, parent_id, comment) values(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, comment.TaskId, comment.UserId, comment.ParentId, comment.Comment)
	helper.PanifIfError(err)

	id, err := result.LastInsertId()
	helper.PanifIfError(err)

	comment.Id = int(id)
	return comment
}

func (repository *CommentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	SQL := "UPDATE comments SET comment = ?, WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, comment.Comment, comment.Id)
	helper.PanifIfError(err)

	return comment
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, comment domain.Comment) {
	SQL := "DELETE FROM comments WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, comment.Id)
	helper.PanifIfError(err)
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.Comment, error) {
	SQL := "SELECT id, comment, created_at FROM comments WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, commentId)
	helper.PanifIfError(err)
	defer rows.Close()

	comment := domain.Comment{}
	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Comment, &comment.CreatedAt)
		helper.PanifIfError(err)
		return comment, nil
	} else {
		return comment, errors.New("Comment Is Not Found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Comment {
	SQL := "SELECT id, task_id, user_id, parent_id, comment, created_at FROM comments"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanifIfError(err)
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		comment := domain.Comment{}
		err := rows.Scan(&comment.Id, &comment.TaskId, &comment.UserId, &comment.ParentId, &comment.Comment, &comment.CreatedAt)
		helper.PanifIfError(err)
		comments = append(comments, comment)
	}

	return comments
}
