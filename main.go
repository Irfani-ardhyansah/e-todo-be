package main

import (
	"e-todo/app"
	"e-todo/controller"
	"e-todo/helper"
	"e-todo/middleware"
	"e-todo/repository"
	"e-todo/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	db := app.NewDB()
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

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	fmt.Println("SERVER IS RUNNING ON PORT 3000")
	helper.PanifIfError(err)
}
