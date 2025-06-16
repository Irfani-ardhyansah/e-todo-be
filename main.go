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
	"github.com/rs/cors"
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

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	commentRepository := repository.NewCommentRepository()
	commentService := service.NewCommentService(commentRepository, db, validate)
	commentController := controller.NewCommentController(commentService)

	authRespository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRespository, userRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router := app.NewRouter(taskController, timerController, timerHistoryController, userController, authController, commentController)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Change this to your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Api-Key"},
		AllowCredentials: true,
	})

	corsRouter := corsHandler.Handler(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(corsRouter),
	}

	err := server.ListenAndServe()
	fmt.Println("SERVER IS RUNNING ON PORT 3000")
	helper.PanifIfError(err)
}
