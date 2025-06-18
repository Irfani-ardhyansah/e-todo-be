package controller

import (
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (controller *CommentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentCreateRequest := web.CommentCreateRequest{}
	helper.ReadFromRequestBody(request, &commentCreateRequest)

	taskResponse := controller.CommentService.Create(request.Context(), commentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentUpdateRequest := web.CommentUpdateRequest{}
	helper.ReadFromRequestBody(request, &commentUpdateRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	commentUpdateRequest.Id = id
	commentResponse := controller.CommentService.Update(request.Context(), commentUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   commentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	controller.CommentService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	taskResponse := controller.CommentService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanifIfError(err)

	taskResponses := controller.CommentService.FindAll(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponses,
	}

	userClaims := helper.UserClaims(request)

	fmt.Println(userClaims["id"])

	helper.WriteToResponseBody(writer, webResponse)
}
