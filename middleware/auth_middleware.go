package middleware

import (
	exception "e-todo/excception"
	"errors"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers,X-Api-Key")
	xApiKey := request.Header.Get("X-Api-Key")
	if len(xApiKey) != 0 && xApiKey != "RAHASISA" {
		err := errors.New("x-api-key Is Not Match")
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	middleware.Handler.ServeHTTP(writer, request)
}
