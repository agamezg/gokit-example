package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}

func MakeEndpoints(service Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(service),
		GetUser: makeGetUserEndpoint(service),
	}
}

func makeCreateUserEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		ok, err := service.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(service Service) endpoint.Endpoint {
 return func(ctx context.Context, request interface{}) (response interface{}, err error) {
 	req := request.(GetUserRequest)
 	email, err := service.GetUser(ctx, req.Id)
 	return GetUserResponse{Email: email}, err
 }
}
