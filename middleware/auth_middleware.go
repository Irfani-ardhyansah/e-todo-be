package middleware

import (
	"e-todo/helper"
	"e-todo/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers,X-Api-Key")
	// for k, v := range request.Header {
	// 	fmt.Printf("Header: %s = %s\n", k, v)
	// }
	// origin := request.Header.Get("Origin")
	// fmt.Println("origin ", origin)
	// fmt.Println("x-api-key", request.Header.Get("X-Api-Key"))
	xApiKey := request.Header.Get("X-Api-Key")
	// fmt.Println(len(xApiKey))
	if len(xApiKey) != 0 && xApiKey != "RAHASISA" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		// fmt.Println("NOT")
	}

	middleware.Handler.ServeHTTP(writer, request)
}
