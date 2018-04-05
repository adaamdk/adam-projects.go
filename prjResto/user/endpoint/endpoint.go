package endpoint

import (
	"context"

	svc "prjResto/user/server"

	kit "github.com/go-kit/kit/endpoint"
)

// 5.
type UserEndpoint struct {
	AddUserEndpoint      kit.Endpoint
	UpdateUserEndpoint   kit.Endpoint
	ReadUserEndpoint     kit.Endpoint
	ReadUserByIDEndpoint kit.Endpoint
}

func NewUserEndpoint(service svc.UserService) UserEndpoint {
	addUserEp := makeAddUserEndpoint(service)
	updateUserEp := makeUpdateUserEndpoint(service)
	readUserEp := makeReadUserEndpoint(service)
	readUserByIDEp := makeReadUserByIDEndpoint(service)

	return UserEndpoint{AddUserEndpoint: addUserEp,
		UpdateUserEndpoint:   updateUserEp,
		ReadUserEndpoint:     readUserEp,
		ReadUserByIDEndpoint: readUserByIDEp,
	}
}

func makeAddUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		err := service.AddUserService(ctx, req)
		return nil, err
	}
}

func makeUpdateUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ct context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		err := service.UpdateUserService(ct, req)
		return nil, err
	}
}

func makeReadUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ct context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadUserService(ct)
		return result, err
	}
}
func makeReadUserByIDEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		result, err := service.ReadUserByIDService(ctx, req.ID)

		return result, err
	}
}
