package main

import (
	"fmt"

	ep "prjResto/gudang/endpoint"
	pb "prjResto/gudang/grpc"
	svc "prjResto/gudang/server"

	cfg "prjResto/util/config"
	run "prjResto/util/grpc"
	util "prjResto/util/microservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// import dari utility
	// logging
	logger := util.Logger()

	ok := cfg.AppConfig.LoadConfig()
	if !ok {
		logger.Log(util.LogError, "failed to load configuration")
		return
	}
	// diambil dari service.conf
	discHost := cfg.GetA("discoveryhost", "127.0.0.1:2181")
	ip := cfg.Get("serviceip", "127.0.0.1")
	port := cfg.Get("serviceport", "7001")
	address := fmt.Sprintf("%s:%s", ip, port)

	registrar, err := util.ServiceRegistry(discHost, svc.ServiceID, address, logger)
	if err != nil {
		logger.Log(util.LogError, "cannot find register")
		return
	}
	registrar.Register()
	defer registrar.Deregister()

	tracerHost := cfg.Get("tracerhost", "127.0.0.1:9999")
	tracer := util.Tracer(tracerHost)
	// 1.
	var server pb.GudangServiceServer
	{

		dbHost := cfg.Get(cfg.DBhost, "127.0.0.1:3306")
		dbName := cfg.Get(cfg.DBname, "Restoran")
		dbGudang := cfg.Get(cfg.DBuid, "root")
		dbPwd := cfg.Get(cfg.DBpwd, "root")

		//brokers := cfg.GetA("mqbrokers", "127.0.0.1:9092")

		//before code
		dbReadWriter := svc.NewDBReadWriter(dbHost, dbName, dbGudang, dbPwd)
		//dbRuler := svc.NewDBDispatchRuler(dbReadWriter, locator)
		//notifier := mq.NewAsyncProducer(brokers, nil)

		//auctioneer := svc.NewAuctioneer(dbReadWriter, cacher)
		service := svc.NewGudang(dbReadWriter)
		endpoint := ep.NewGudangEndpoint(service)
		fmt.Println(endpoint)
		server = ep.NewGRPCGudangServer(endpoint, tracer, logger)
	}

	grpcServer := grpc.NewServer(run.Recovery(logger)...)
	pb.RegisterGudangServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	exit := make(chan bool, 1)
	go run.Serve(address, grpcServer, exit, logger)

	<-exit
}
