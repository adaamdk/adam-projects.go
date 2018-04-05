package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "jadwalshift.bluebird.id"
	OnAdd     Status = 1
)

type JadwalShift struct {
	Code      string
	JadwalOn  string
	JadwalOff string
	Status    int32
	CreatedBy string
	CreatedOn string
	UpdatedBy string
	UpdatedOn string
}

// utk parameter
type ReadWriter interface {
	AddJadwalShift(JadwalShift) error
	UpdateJadwalShift(JadwalShift) error
}

// utk fungsi nilai return
type JadwalShiftService interface {
	AddJadwalShiftService(context.Context, JadwalShift) error
	UpdateJadwalShiftService(context.Context, JadwalShift) error
}
