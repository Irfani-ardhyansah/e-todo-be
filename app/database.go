package app

import (
	"database/sql"
	"e-todo/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:hwhwhwlol@tcp(localhost:33061)/db_study_todo")
	helper.PanifIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10)

	return db
}
