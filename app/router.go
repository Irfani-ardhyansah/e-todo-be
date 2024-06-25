package app

import (
	"e-todo/controller"
	exception "e-todo/excception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(taskController controller.TaskController, timerController controller.TimerController, timerHistoryController controller.TimerHistoryController) *httprouter.Router {
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

	return router
}
