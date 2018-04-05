package endpoint

import (
	"context"

	svc "prjResto/gudang/server"

	kit "github.com/go-kit/kit/endpoint"
)

// 5.
type GudangEndpoint struct {
	AddGudangEndpoint              kit.Endpoint
	UpdateGudangEndpoint           kit.Endpoint
	ReadGudangByKeteranganEndpoint kit.Endpoint
}

func NewGudangEndpoint(service svc.GudangService) GudangEndpoint {
	addGudangEp := makeAddGudangEndpoint(service)
	updateGudangEp := makeUpdateGudangEndpoint(service)
	readGudangByKeteranganEp := makeReadGudangByKeteranganEndpoint(service)

	return GudangEndpoint{AddGudangEndpoint: addGudangEp,
		UpdateGudangEndpoint:           updateGudangEp,
		ReadGudangByKeteranganEndpoint: readGudangByKeteranganEp,
	}
}

func makeAddGudangEndpoint(service svc.GudangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Gudang)
		err := service.AddGudangService(ctx, req)
		return nil, err
	}
}

func makeUpdateGudangEndpoint(service svc.GudangService) kit.Endpoint {
	return func(ct context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Gudang)
		err := service.UpdateGudangService(ct, req)
		return nil, err
	}
}

func makeReadGudangByKeteranganEndpoint(service svc.GudangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Gudang)
		result, err := service.ReadGudangByKeteranganService(ctx, req.Keterangan)

		return result, err
	}
}
