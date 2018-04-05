package server

import (
	"context"
	"fmt"
)

// 5.
type jadwalShift struct {
	writer ReadWriter // diambil dari interface di service.go
}

func NewJadwalShift(writer ReadWriter) JadwalShiftService {
	return &jadwalShift{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (js *jadwalShift) AddJadwalShiftService(ctx context.Context, jadwalShift JadwalShift) error {
	fmt.Println("Input data berhasil, silahkan periksa DB Anda, gan!")
	err := js.writer.AddJadwalShift(jadwalShift)
	if err != nil {
		return err
	}
	return nil
}

func (js *jadwalShift) UpdateJadwalShiftService(ctx context.Context, sh JadwalShift) error {
	err := js.writer.UpdateJadwalShift(sh)
	if err != nil {
		return err
	}
	return nil
}
