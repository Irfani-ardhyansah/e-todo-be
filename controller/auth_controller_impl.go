package controller

import (
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &authCreateRequest)

	authResponse := controller.AuthService.Login(request.Context(), authCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) RefreshToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	refreshTokenRequest := web.RefreshTokenRequest{}
	helper.ReadFromRequestBody(request, &refreshTokenRequest)

	refreshTokenResponse := controller.AuthService.RefreshToken(request.Context(), refreshTokenRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   refreshTokenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
