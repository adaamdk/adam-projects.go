package main

import (
	"context"
	"fmt"
	"time"

	cli "prjResto/user/endpoint"
	opt "prjResto/util/grpc"
	util "prjResto/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCUserClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add
	//client.AddUserService(context.Background(), svc.User{ID: "003", Username: "raniyus", Pwd: "r4n1y", IDkaryawan: "K001", CreatedBy: "Adam", UpdatedBy: "Adam"})

	//Update
	//client.UpdateUserService(context.Background(), svc.User{ID: "003", Username: "Rafi", Pwd: "s1An*aY", IDkaryawan: "K001", CreatedBy: "Adam", UpdatedBy: "Adam"})

	//select
	//parameter := "1%"
	id, _ := client.ReadUserByIDService(context.Background(), "001")
	//, parameter)
	fmt.Println("ID :", id)

	//select all
	//client.ReadUserService(context.Background())
}
