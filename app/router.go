package app

import (
	"context"
	"e-todo/controller"
	exception "e-todo/excception"
	"e-todo/helper"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func jwtMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			err := errors.New("Missing Authorization header")
			panic(exception.NewUnauthorizedError(err.Error()))
		}

		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			err := errors.New("Invalid Authorization header format")
			panic(exception.NewUnauthorizedError(err.Error()))
		}
		token := tokenParts[1]

		claims, err := helper.VerifyToken(token, "access")
		if err != nil {
			panic(exception.NewUnauthorizedError(err.Error()))
		}

		ctx := context.WithValue(r.Context(), "jwtClaims", claims)

		next(w, r.WithContext(ctx), ps)
	}
}

func NewRouter(taskController controller.TaskController, timerController controller.TimerController, timerHistoryController controller.TimerHistoryController, userController controller.UserController, authController controller.AuthController, commentController controller.CommentController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/v1/tasks", jwtMiddleware(taskController.FindAll))
	router.GET("/api/v1/task/:taskId", jwtMiddleware(taskController.FindById))
	router.POST("/api/v1/task", jwtMiddleware(taskController.Create))
	router.PUT("/api/v1/task/:taskId", jwtMiddleware(taskController.Update))
	router.PUT("/api/v1/task/:taskId/status", jwtMiddleware(taskController.UpdateStatus))
	router.DELETE("/api/v1/task/:taskId", jwtMiddleware(taskController.Delete))

	router.GET("/api/v1/task/:taskId/comments", jwtMiddleware(commentController.FindAll))
	router.GET("/api/v1/task/:taskId/comments/:commentId", jwtMiddleware(commentController.FindById))
	router.POST("/api/v1/task/:taskId/comments", jwtMiddleware(commentController.Create))
	router.PUT("/api/v1/task/:taskId/comments/:commentId", jwtMiddleware(commentController.Update))
	router.DELETE("/api/v1/task/:taskId/comments", jwtMiddleware(commentController.Delete))

	router.POST("/api/v1/timer/start/:taskId", jwtMiddleware(timerController.Create))
	router.PUT("/api/v1/timer/update/:timerId", jwtMiddleware(timerController.Update))

	router.GET("/api/v1/timer/history/:timerId", jwtMiddleware(timerHistoryController.ListByTimer))
	router.GET("/api/v1/timer/weekly-report", jwtMiddleware(timerHistoryController.ListWeeklyReport))
	router.GET("/api/v1/timer/weekly-report/:taskId", jwtMiddleware(timerHistoryController.ListWeeklyReportByTaskId))

	router.POST("/api/v1/user/create", jwtMiddleware(userController.Create))

	router.POST("/api/v1/user/login", authController.Login)
	router.POST("/api/v1/user/refresh-token", authController.RefreshToken)

	router.PanicHandler = exception.ErrorHandler

	return router
}
