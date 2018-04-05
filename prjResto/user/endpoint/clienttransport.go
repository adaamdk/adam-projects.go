package endpoint

import (
	"context"
	"time"

	pb "prjResto/user/grpc"
	svc "prjResto/user/server"
	util "prjResto/util/grpc"
	disc "prjResto/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.UserService"
)

func NewGRPCUserClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.UserService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddUserEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addUserEp = retry
	}
	var updateUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateUser, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateUserEp = retry
	}

	var readUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadUserEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readUserEp = retry
	}

	var readUserByIDEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadUserByIDEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readUserByIDEp = retry
	}

	return UserEndpoint{AddUserEndpoint: addUserEp,
		UpdateUserEndpoint:   updateUserEp,
		ReadUserEndpoint:     readUserEp,
		ReadUserByIDEndpoint: readUserByIDEp}, nil
}

func encodeAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.AddUserReq{
		ID:         req.ID,
		Username:   req.Username,
		Pwd:        req.Pwd,
		IDkaryawan: req.IDkaryawan,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		UpdatedBy:  req.UpdatedBy,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.UpdateUserReq{
		ID:         req.ID,
		Username:   req.Username,
		Pwd:        req.Pwd,
		IDkaryawan: req.IDkaryawan,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdatedOn:  req.UpdatedOn,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeReadUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadUserByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.ReadUserByIDReq{ID: req.ID}, nil
}

func decodeUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadUserByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadUserByIDResp)
	return svc.User{
		ID:         resp.ID,
		Username:   resp.Username,
		Pwd:        resp.Pwd,
		IDkaryawan: resp.IDkaryawan,
		Status:     resp.Status,
		CreatedBy:  resp.CreatedBy,
		CreatedOn:  resp.CreatedOn,
		UpdatedBy:  resp.UpdatedBy,
		UpdatedOn:  resp.UpdatedOn,
		Keterangan: resp.Keterangan,
	}, nil
}

func decodeReadUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadUserResp)
	var rsp svc.Users

	for _, r := range resp.AllUser {
		itm := svc.User{
			ID:         r.ID,
			Username:   r.Username,
			Pwd:        r.Pwd,
			IDkaryawan: r.IDkaryawan,
			Status:     r.Status,
			CreatedBy:  r.CreatedBy,
			CreatedOn:  r.CreatedOn,
			UpdatedBy:  r.UpdatedBy,
			UpdatedOn:  r.UpdatedOn,
			Keterangan: r.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddUserEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddUser",
		encodeAddUserRequest,
		decodeUserResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateUser(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateUser",
		encodeUpdateUserRequest,
		decodeUserResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadUserEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadUser",
		encodeReadUserRequest,
		decodeReadUserResponse,
		pb.ReadUserResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
func makeClientReadUserByIDEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadUserByID",
		encodeReadUserByIDRequest,
		decodeReadUserByIDResponse,
		pb.ReadUserByIDResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadUserByID")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadUserByID",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
