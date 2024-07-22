package app

import (
	"context"
	"e-todo/controller"
	exception "e-todo/excception"
	"e-todo/helper"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func jwtMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]

		claims, err := helper.VerifyToken(token, "access")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "jwtClaims", claims)

		next(w, r.WithContext(ctx), ps)
	}
}

func NewRouter(taskController controller.TaskController, timerController controller.TimerController, timerHistoryController controller.TimerHistoryController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/v1/tasks", jwtMiddleware(taskController.FindAll))
	router.GET("/api/v1/task/:taskId", taskController.FindById)
	router.POST("/api/v1/task", taskController.Create)
	router.PUT("/api/v1/task/:taskId", taskController.Update)
	router.PUT("/api/v1/task-status/:taskId", taskController.UpdateStatus)
	router.DELETE("/api/v1/task/:taskId", taskController.Delete)

	router.POST("/api/v1/timer/start/:taskId", timerController.Create)
	router.PUT("/api/v1/timer/update/:timerId", timerController.Update)

	router.GET("/api/v1/timer/history/:timerId", timerHistoryController.ListByTimer)

	router.POST("/api/v1/user/create", userController.Create)

	router.POST("/api/v1/user/login", userController.Login)

	router.POST("/api/v1/user/refresh-token/:userId", userController.RefreshToken)

	router.PanicHandler = exception.ErrorHandler

	return router
}
