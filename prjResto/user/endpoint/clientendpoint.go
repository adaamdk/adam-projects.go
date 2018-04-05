package endpoint

import (
	"context"
	f "fmt"

	svc "prjResto/user/server"
)

func (ue UserEndpoint) AddUserService(ctx context.Context, usr svc.User) error {
	_, err := ue.AddUserEndpoint(ctx, usr)
	return err
}

func (ue UserEndpoint) UpdateUserService(ctx context.Context, usr svc.User) error {
	_, err := ue.UpdateUserEndpoint(ctx, usr)
	if err != nil {
		f.Println("error pada endpoint:", err)
	}
	return err
}

func (ue UserEndpoint) ReadUserService(ctx context.Context) (svc.Users, error) {
	response, err := ue.ReadUserEndpoint(ctx, nil)
	f.Println("Response", response)
	if err != nil {
		f.Println("Error di endpoint lu gan..", err)
	}
	return response.(svc.Users), err
}

func (ue UserEndpoint) ReadUserByIDService(ctx context.Context, ID string) (svc.User, error) {
	req := svc.User{ID: ID}
	f.Println(req)
	resp, err := ue.ReadUserByIDEndpoint(ctx, req)
	if err != nil {
		f.Println("error pada endpoint: ", err)
	}
	us := resp.(svc.User)
	return us, err
}
