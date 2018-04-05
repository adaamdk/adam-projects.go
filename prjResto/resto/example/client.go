package main

import (
	"context"
	"time"

	cli "prjResto/resto/endpoint"
	svc "prjResto/resto/server"
	opt "prjResto/util/grpc"
	util "prjResto/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCJadwalShiftClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add
	//client.AddCustomerService(context.Background(), svc.Customer{Name: "jkl", CustomerType: 1, Mobile: "087889", Email: "jkl@gmail", Gender: "F", CallbackPhone: "087889"})
	client.AddJadwalShiftService(context.Background(), svc.JadwalShift{Code: "SH006", JadwalOn: "2018-03-20", JadwalOff: "2018-03-21", CreatedBy: "Adam", UpdatedBy: "Adam"})

	//Update
	client.UpdateJadwalShiftService(context.Background(), svc.JadwalShift{Code: "SH003", JadwalOn: "2018-03-13", JadwalOff: "2018-03-14", Status: 1, CreatedBy: "Luffy", UpdatedBy: "Adam"})
	/*
		//Get Customer By Mobile No
		cusMobile, _ := client.ReadCustomerByMobileService(context.Background(), "087889")
		fmt.Println("customer based on mobile:",cusMobile)

		//List Customer
		cuss, _ := client.ReadCustomerService(context.Background())
		fmt.Println("all customers:",cuss)

	*/
}
