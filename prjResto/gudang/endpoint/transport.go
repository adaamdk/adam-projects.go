package endpoint

import (
	"context"

	pb "prjResto/gudang/grpc"
	scv "prjResto/gudang/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcGudangServer struct {
	addGudang              grpctransport.Handler
	updateGudang           grpctransport.Handler
	readGudangByKeterangan grpctransport.Handler
}

func NewGRPCGudangServer(endpoints GudangEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.GudangServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcGudangServer{
		addGudang: grpctransport.NewServer(endpoints.AddGudangEndpoint,
			decodeAddGudangRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddGudang", logger)))...),

		updateGudang: grpctransport.NewServer(endpoints.UpdateGudangEndpoint,
			decodeUpdateGudangRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateGudang", logger)))...),

		readGudangByKeterangan: grpctransport.NewServer(endpoints.ReadGudangByKeteranganEndpoint,
			decodeReadGudangByKeteranganRequest,
			encodeReadGudangByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadGudangByKeterangan", logger)))...),
	}
}

func decodeAddGudangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddGudangReq)
	return scv.Gudang{
		ID:         req.GetID(),
		Name:       req.GetNamaGudang(),
		Alamat:     req.GetAlamatGudang(),
		Luas:       req.GetLuasGudang(),
		Status:     req.GetStatus(),
		CreatedBy:  req.GetCreatedBy(),
		CreatedOn:  req.GetCreatedOn(),
		UpdatedBy:  req.GetUpdatedBy(),
		UpdatedOn:  req.GetUpdatedOn(),
		Keterangan: req.GetKeterangan()}, nil
}

func decodeUpdateGudangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateGudangReq)
	return scv.Gudang{
		ID:         req.ID,
		Name:       req.NamaGudang,
		Alamat:     req.AlamatGudang,
		Luas:       req.LuasGudang,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdatedOn:  req.UpdatedOn,
		Keterangan: req.Keterangan}, nil
}

func encodeReadGudangByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	r := response.(scv.Gudangs)

	rsp := &pb.ReadGudangByKeteranganResp{}
	for _, u := range r {
		gd := &pb.ReadGudangByKeterangan{
			ID:         u.ID,
			NamaGudang: u.Name,
			Keterangan: u.Keterangan,
		}
		rsp.AllGudang = append(rsp.AllGudang, gd)
	}
	return r, nil
}

func decodeReadGudangByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadGudangByKeteranganReq)
	return scv.Gudang{Keterangan: req.Keterangan}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcGudangServer) AddGudang(ctx oldcontext.Context, shift *pb.AddGudangReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addGudang.ServeGRPC(ctx, shift)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcGudangServer) UpdateGudang(ctx oldcontext.Context, js *pb.UpdateGudangReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateGudang.ServeGRPC(ctx, js)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcGudangServer) ReadGudangByKeterangan(ctx oldcontext.Context, e *pb.ReadGudangByKeteranganReq) (*pb.ReadGudangByKeteranganResp, error) {
	_, resp, err := s.readGudangByKeterangan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadGudangByKeteranganResp), nil
}
