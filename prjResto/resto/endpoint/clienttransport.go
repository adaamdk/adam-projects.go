package endpoint

import (
	"context"
	"time"

	pb "prjResto/resto/grpc"
	svc "prjResto/resto/server"
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
	grpcName = "grpc.JadwalShiftService"
)

func NewGRPCJadwalShiftClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.JadwalShiftService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addJadwalShiftEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddJadwalShiftEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addJadwalShiftEp = retry
	}
	var updateJadwalShiftEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateJadwalShift, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateJadwalShiftEp = retry
	}

	return JadwalShiftEndpoint{AddJadwalShiftEndpoint: addJadwalShiftEp,
		UpdateJadwalShiftEndpoint: updateJadwalShiftEp}, nil
}

func encodeAddJadwalShiftRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.JadwalShift)
	return &pb.AddJadwalShiftReq{
		ShiftCode:    req.Code,
		JadwalMasuk:  req.JadwalOn,
		JadwalPulang: req.JadwalOff,
		Status:       req.Status,
		CreatedBy:    req.CreatedBy,
		UpdatedBy:    req.UpdatedBy,
	}, nil
}

func encodeUpdateJadwalShiftRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.JadwalShift)
	return &pb.UpdateJadwalShiftReq{
		ShiftCode:    req.Code,
		JadwalMasuk:  req.JadwalOn,
		JadwalPulang: req.JadwalOff,
		Status:       req.Status,
		CreatedBy:    req.CreatedBy,
		CreatedOn:    req.CreatedOn,
		UpdatedBy:    req.UpdatedBy,
		UpdatedOn:    req.UpdatedOn,
	}, nil
}

func decodeJadwalShiftResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func makeClientAddJadwalShiftEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddJadwalShift",
		encodeAddJadwalShiftRequest,
		decodeJadwalShiftResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddJadwalShift")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddJadwalShift",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateJadwalShift(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateJadwalShift",
		encodeUpdateJadwalShiftRequest,
		decodeJadwalShiftResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateJadwalShift")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateJadwalShift",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
