package account

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}

	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeEmailRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(request)

	req = GetUserRequest{
		Id: vars["id"],
	}
	
	return req, nil
}
