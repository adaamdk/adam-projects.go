package server

import "context"

type Status int32

// 1.
const (
	//ServiceID is dispatch service ID
	ServiceID        = "daftaruser.bluebird.id"
	OnAdd     Status = 1
)

type User struct {
	ID         string
	Username   string
	Pwd        string
	IDkaryawan string
	Status     int32
	CreatedBy  string
	CreatedOn  string
	UpdatedBy  string
	UpdatedOn  string
	Keterangan string
}
type Users []User

// utk parameter
type ReadWriter interface {
	AddUser(User) error
	UpdateUser(User) error
	ReadUser() (Users, error)
	ReadUserByID(string) (User, error)
}

// utk fungsi nilai return
type UserService interface {
	AddUserService(context.Context, User) error
	UpdateUserService(context.Context, User) error
	ReadUserService(context.Context) (Users, error)
	ReadUserByIDService(context.Context, string) (User, error)
}
