package account

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailRequest,
		encodeResponse,
	))

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
