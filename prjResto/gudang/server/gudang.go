package server

import (
	"context"
	"fmt"
)

// 4.
type gudang struct {
	writer ReadWriter // diambil dari interface di service.go
}

func NewGudang(writer ReadWriter) GudangService {
	return &gudang{writer: writer}
}

//Methode pada interface GudangService di service.go
func (g *gudang) AddGudangService(ctx context.Context, gudang Gudang) error {
	fmt.Println("Input data berhasil, silahkan periksa DB Anda, gan!")
	err := g.writer.AddGudang(gudang)
	if err != nil {
		return err
	}
	return nil
}

func (g *gudang) UpdateGudangService(ctx context.Context, gd Gudang) error {
	err := g.writer.UpdateGudang(gd)
	if err != nil {
		return err
	}
	return nil
}

func (g *gudang) ReadGudangByKeteranganService(ctx context.Context, k string) (Gudangs, error) {
	us, err := g.writer.ReadGudangByKeterangan(k)
	fmt.Println("Keterangan", us)
	if err != nil {
		return us, err
	}
	return us, nil
}
