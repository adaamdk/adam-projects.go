package server

import (
	"context"
	"fmt"
)

// 5.
type user struct {
	writer ReadWriter // diambil dari interface di service.go
}

func NewUser(writer ReadWriter) UserService {
	return &user{writer: writer}
}

//Methode pada interface UserService di service.go
func (u *user) AddUserService(ctx context.Context, user User) error {
	fmt.Println("Input data berhasil, silahkan periksa DB Anda, gan!")
	err := u.writer.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) UpdateUserService(ctx context.Context, usr User) error {
	err := u.writer.UpdateUser(usr)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) ReadUserService(ctx context.Context) (Users, error) {
	us, err := u.writer.ReadUser()
	//fmt.Println("User", us)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (u *user) ReadUserByIDService(ctx context.Context, s string) (User, error) {
	usr, err := u.writer.ReadUserByID(s)
	//fmt.Println("User", usr)
	if err != nil {
		return usr, err
	}
	return usr, nil
}
