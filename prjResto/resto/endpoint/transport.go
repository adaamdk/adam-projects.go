package endpoint

import (
	"context"

	pb "prjResto/resto/grpc"
	scv "prjResto/resto/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcJadwalShiftServer struct {
	addJadwalShift    grpctransport.Handler
	updateJadwalShift grpctransport.Handler
}

func NewGRPCJadwalShiftServer(endpoints JadwalShiftEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.JadwalShiftServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcJadwalShiftServer{
		addJadwalShift: grpctransport.NewServer(endpoints.AddJadwalShiftEndpoint,
			decodeAddJadwalShiftRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddJadwalShift", logger)))...),

		updateJadwalShift: grpctransport.NewServer(endpoints.UpdateJadwalShiftEndpoint,
			decodeUpdateJadwalShiftRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateJadwalShift", logger)))...),
	}
}

func decodeAddJadwalShiftRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddJadwalShiftReq)
	return scv.JadwalShift{
		Code:      req.GetShiftCode(),
		JadwalOn:  req.GetJadwalMasuk(),
		JadwalOff: req.GetJadwalPulang(),
		Status:    req.GetStatus(),
		CreatedBy: req.GetCreatedBy(),
		CreatedOn: req.GetCreatedOn(),
		UpdatedBy: req.GetUpdatedBy(),
		UpdatedOn: req.GetUpdatedOn()}, nil
}

func decodeUpdateJadwalShiftRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJadwalShiftReq)
	return scv.JadwalShift{
		Code:      req.ShiftCode,
		JadwalOn:  req.JadwalMasuk,
		JadwalOff: req.JadwalPulang,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		CreatedOn: req.CreatedOn,
		UpdatedBy: req.UpdatedBy,
		UpdatedOn: req.UpdatedOn}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcJadwalShiftServer) AddJadwalShift(ctx oldcontext.Context, shift *pb.AddJadwalShiftReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addJadwalShift.ServeGRPC(ctx, shift)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcJadwalShiftServer) UpdateJadwalShift(ctx oldcontext.Context, js *pb.UpdateJadwalShiftReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateJadwalShift.ServeGRPC(ctx, js)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
