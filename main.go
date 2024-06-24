package main

import (
	"e-todo/app"
	"e-todo/controller"
	exception "e-todo/excception"
	"e-todo/helper"
	"e-todo/middleware"
	"e-todo/repository"
	"e-todo/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
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

	router := httprouter.New()
	router.GET("/api/v1/tasks", taskController.FindAll)
	router.GET("/api/v1/task/:taskId", taskController.FindById)
	router.POST("/api/v1/task", taskController.Create)
	router.PUT("/api/v1/task/:taskId", taskController.Update)
	router.PUT("/api/v1/task-status/:taskId", taskController.UpdateStatus)
	router.DELETE("/api/v1/task/:taskId", taskController.Delete)

	router.POST("/api/v1/timer/start/:taskId", timerController.Create)
	router.PUT("/api/v1/timer/update/:timerId", timerController.Update)

	router.GET("/api/v1/timer/history/:timerId", timerHistoryController.ListByTimer)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	fmt.Println("SERVER IS RUNNING ON PORT 3000")
	helper.PanifIfError(err)
}
