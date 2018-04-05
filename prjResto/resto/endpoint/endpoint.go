package endpoint

import (
	"context"

	svc "prjResto/resto/server"

	kit "github.com/go-kit/kit/endpoint"
)

type JadwalShiftEndpoint struct {
	AddJadwalShiftEndpoint    kit.Endpoint
	UpdateJadwalShiftEndpoint kit.Endpoint
}

func NewJadwalShiftEndpoint(service svc.JadwalShiftService) JadwalShiftEndpoint {
	addJadwalShiftEp := makeAddJadwalShiftEndpoint(service)
	updateJadwalShiftEp := makeUpdateJadwalShiftEndpoint(service)

	return JadwalShiftEndpoint{AddJadwalShiftEndpoint: addJadwalShiftEp,
		UpdateJadwalShiftEndpoint: updateJadwalShiftEp,
	}
}

func makeAddJadwalShiftEndpoint(service svc.JadwalShiftService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.JadwalShift)
		err := service.AddJadwalShiftService(ctx, req)
		return nil, err
	}
}

func makeUpdateJadwalShiftEndpoint(service svc.JadwalShiftService) kit.Endpoint {
	return func(ct context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.JadwalShift)
		err := service.UpdateJadwalShiftService(ct, req)
		return nil, err
	}
}
