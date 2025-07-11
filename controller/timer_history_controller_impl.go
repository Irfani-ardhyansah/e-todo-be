package controller

import (
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TimerHistoryControllerImpl struct {
	TimerHistoryService service.TimerHistoryService
}

func NewTimerHistoryController(timerHistoryService service.TimerHistoryService) TimerHistoryController {
	return &TimerHistoryControllerImpl{
		TimerHistoryService: timerHistoryService,
	}
}

func (controller *TimerHistoryControllerImpl) ListByTimer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	timerId := params.ByName("timerId")
	id, err := strconv.Atoi(timerId)
	helper.PanifIfError(err)

	timerHistoryResponse := controller.TimerHistoryService.FindByParentId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   timerHistoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TimerHistoryControllerImpl) ListWeeklyReport(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	weeklyReportResponse := controller.TimerHistoryService.GetWeeklyReport(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   weeklyReportResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TimerHistoryControllerImpl) ListWeeklyReportByTaskId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	weeklyReportResponse := controller.TimerHistoryService.GetWeeklyReportByTaskId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   weeklyReportResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
