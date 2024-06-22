package controller

import (
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TaskControllerImpl struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return &TaskControllerImpl{
		TaskService: taskService,
	}
}

func (controller *TaskControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskCreateRequest := web.TaskCreateRequest{}
	helper.ReadFromRequestBody(request, &taskCreateRequest)

	taskResponse := controller.TaskService.Create(request.Context(), taskCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TaskControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskUpdateRequest := web.TaskUpdateRequest{}
	helper.ReadFromRequestBody(request, &taskUpdateRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	taskUpdateRequest.Id = id
	taskResponse := controller.TaskService.Update(request.Context(), taskUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TaskControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskUpdateRequest := web.TaskUpdateStatusRequest{}
	helper.ReadFromRequestBody(request, &taskUpdateRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	taskUpdateRequest.Id = id
	controller.TaskService.UpdateStatus(request.Context(), taskUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TaskControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	controller.TaskService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TaskControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	taskResponse := controller.TaskService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TaskControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskResponses := controller.TaskService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
