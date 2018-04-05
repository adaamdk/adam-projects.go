package server

import "context"

type Status int32

// 1.
const (
	//ServiceID is dispatch service ID
	ServiceID        = "gudang.bluebird.id"
	OnAdd     Status = 1
)

type Gudang struct {
	ID         string
	Name       string
	Alamat     string
	Luas       string
	Status     int32
	CreatedBy  string
	CreatedOn  string
	UpdatedBy  string
	UpdatedOn  string
	Keterangan string
}

type Gudangs []Gudang

// utk parameter
type ReadWriter interface {
	AddGudang(Gudang) error
	UpdateGudang(Gudang) error
	ReadGudangByKeterangan(string) (Gudangs, error)
}

// utk fungsi nilai return
type GudangService interface {
	AddGudangService(context.Context, Gudang) error
	UpdateGudangService(context.Context, Gudang) error
	ReadGudangByKeteranganService(context.Context, string) (Gudangs, error)
}
