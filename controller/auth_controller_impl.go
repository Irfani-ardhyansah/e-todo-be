package controller

import (
	"e-todo/helper"
	"e-todo/model/web"
	"e-todo/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanifIfError(err)

	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(writer, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenParts := strings.SplitN(authHeader, " ", 2)
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		http.Error(writer, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}
	refreshToken := tokenParts[1]

	refreshTokenResponse := controller.AuthService.RefreshToken(request.Context(), id, refreshToken)
	fmt.Println(refreshTokenResponse)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   refreshTokenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
