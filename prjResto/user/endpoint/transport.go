package endpoint

import (
	"context"

	pb "prjResto/user/grpc"
	scv "prjResto/user/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcUserServer struct {
	addUser      grpctransport.Handler
	updateUser   grpctransport.Handler
	readUser     grpctransport.Handler
	readUserByID grpctransport.Handler
}

func NewGRPCUserServer(endpoints UserEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.UserServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcUserServer{
		addUser: grpctransport.NewServer(endpoints.AddUserEndpoint,
			decodeAddUserRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddUser", logger)))...),

		updateUser: grpctransport.NewServer(endpoints.UpdateUserEndpoint,
			decodeUpdateUserRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUser", logger)))...),

		readUser: grpctransport.NewServer(endpoints.ReadUserEndpoint,
			decodeReadUserRequest,
			encodeReadUserResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadUser", logger)))...),

		readUserByID: grpctransport.NewServer(endpoints.ReadUserByIDEndpoint,
			decodeReadUserByIDRequest,
			encodeReadUserByIDResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadUserByID", logger)))...),
	}
}

// decode Add
func decodeAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddUserReq)
	return scv.User{
		ID:         req.GetID(),
		Username:   req.GetUsername(),
		Pwd:        req.GetPwd(),
		IDkaryawan: req.GetIDkaryawan(),
		Status:     req.GetStatus(),
		CreatedBy:  req.GetCreatedBy(),
		CreatedOn:  req.GetCreatedOn(),
		UpdatedBy:  req.GetUpdatedBy(),
		UpdatedOn:  req.GetUpdatedOn(),
		Keterangan: req.GetKeterangan()}, nil
}

//decode update
func decodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserReq)
	return scv.User{
		ID:         req.ID,
		Username:   req.Username,
		Pwd:        req.Pwd,
		IDkaryawan: req.IDkaryawan,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdatedOn:  req.UpdatedOn,
		Keterangan: req.Keterangan}, nil
}

// decode read
func decodeReadUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

// decode
func decodeReadRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadUserByIDReq)
	return scv.User{ID: req.ID}, nil
}

func decodeReadUserByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadUserByIDReq)
	return scv.User{ID: req.ID}, nil
}

func encodeReadUserByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.User)
	return &pb.ReadUserByIDResp{ID: resp.ID, Username: resp.Username, Pwd: resp.Pwd, IDkaryawan: resp.IDkaryawan,
		Status: resp.Status, CreatedBy: resp.CreatedBy, CreatedOn: resp.CreatedOn, UpdatedBy: resp.UpdatedBy, UpdatedOn: resp.UpdatedOn, Keterangan: resp.Keterangan}, nil
}

func encodeReadUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	r := response.(scv.Users)

	rsp := &pb.ReadUserResp{}
	for _, u := range r {
		usr := &pb.ReadUserByIDResp{
			ID:         u.ID,
			Username:   u.Username,
			Pwd:        u.Pwd,
			IDkaryawan: u.IDkaryawan,
			Status:     u.Status,
			CreatedBy:  u.CreatedBy,
			CreatedOn:  u.CreatedOn,
			UpdatedBy:  u.UpdatedBy,
			UpdatedOn:  u.UpdatedOn,
			Keterangan: u.Keterangan,
		}
		rsp.AllUser = append(rsp.AllUser, usr)
	}
	return r, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcUserServer) AddUser(ctx oldcontext.Context, shift *pb.AddUserReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addUser.ServeGRPC(ctx, shift)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcUserServer) UpdateUser(ctx oldcontext.Context, js *pb.UpdateUserReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, js)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcUserServer) ReadUserByID(ctx oldcontext.Context, nim *pb.ReadUserByIDReq) (*pb.ReadUserByIDResp, error) {
	_, resp, err := s.readUserByID.ServeGRPC(ctx, nim)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadUserByIDResp), nil
}

func (s *grpcUserServer) ReadUser(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadUserResp, error) {
	_, resp, err := s.readUser.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadUserResp), nil
}
