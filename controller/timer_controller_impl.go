package controller

import (
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TimerControllerImpl struct {
	TimerService service.TimerService
}

func NewTimerController(timerService service.TimerService) TimerController {
	return &TimerControllerImpl{
		TimerService: timerService,
	}
}

func (controller *TimerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	timerCreateRequest := web.TimerCreateRequest{}
	helper.ReadFromRequestBody(request, &timerCreateRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	timerCreateRequest.TaskId = id
	timerResponse := controller.TimerService.Start(request.Context(), timerCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   timerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TimerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	timerUpdateRequest := web.TimerUpdateRequest{}
	helper.ReadFromRequestBody(request, &timerUpdateRequest)

	timerId := params.ByName("timerId")
	id, err := strconv.Atoi(timerId)
	helper.PanifIfError(err)

	timerUpdateRequest.Id = id
	timerResponse := controller.TimerService.Update(request.Context(), timerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   timerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
