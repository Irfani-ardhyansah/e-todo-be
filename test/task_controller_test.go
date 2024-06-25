package test

import (
	"context"
	"database/sql"
	"e-todo/app"
	"e-todo/controller"
	"e-todo/helper"
	"e-todo/middleware"
	"e-todo/model/domain"
	"e-todo/repository"
	"e-todo/service"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:hwhwhwlol@tcp(localhost:33061)/db_study_todo_test")
	helper.PanifIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	taskRespository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRespository, db, validate)
	taskController := controller.NewTaskController(taskService)

	timerHistoryRepository := repository.NewTimerHistoryRepository()
	timerHistoryService := service.NewTimerHistoryService(timerHistoryRepository, db, validate)
	timerHistoryController := controller.NewTimerHistoryController(timerHistoryService)

	timerRepository := repository.NewTimerRepository()
	timerService := service.NewTimerService(timerRepository, db, validate, timerHistoryRepository)
	timerController := controller.NewTimerController(timerService)

	router := app.NewRouter(taskController, timerController, timerHistoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateDB(db *sql.DB, dbName string) {
	db.Exec("DELETE FROM " + dbName)

	db.Exec("ALTER TABLE " + dbName + " AUTO_INCREMENT = 1;")
}

func TestCreateTaskSuccess(t *testing.T) {
	db := setupTestDB()
	truncateDB(db, "tasks")
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Test", "status": "open"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/task", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RAHASISA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// assert.Equal(t, "OK", responseBody["status"])
	// assert.Equal(t, "Coba", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateTaskFail(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateDB(db, "tasks")

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/task", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RAHASISA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 500, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 500, int(responseBody["code"].(float64)))

}

func TestUpdateTaskSuccess(t *testing.T) {
	db := setupTestDB()
	truncateDB(db, "tasks")
	router := setupRouter(db)

	tx, _ := db.Begin()
	taskRepository := repository.NewTaskRepository()
	task := taskRepository.Save(context.Background(), tx, domain.Task{
		Name:   "Test",
		Status: "open",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{"name" : "Coba", "status": "close"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/task/"+strconv.Itoa(task.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RAHASISA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
}

func TestUpdateTaskFail(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Coba", "status": "close"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/task/5", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RAHASISA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
}

func TestListTaskSuccess(t *testing.T) {
	db := setupTestDB()
	truncateDB(db, "tasks")
	router := setupRouter(db)

	tx, _ := db.Begin()
	taskRepository := repository.NewTaskRepository()
	taskRepository.Save(context.Background(), tx, domain.Task{
		Name:   "Test",
		Status: "open",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/tasks", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RAHASISA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateDB(db, "tasks")
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/tasks", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
}
