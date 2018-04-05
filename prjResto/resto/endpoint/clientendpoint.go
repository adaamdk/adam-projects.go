package endpoint

import (
	"context"
	f "fmt"

	svc "prjResto/resto/server"
)

func (jse JadwalShiftEndpoint) AddJadwalShiftService(ctx context.Context, shift svc.JadwalShift) error {
	_, err := jse.AddJadwalShiftEndpoint(ctx, shift)
	return err
}

func (jse JadwalShiftEndpoint) UpdateJadwalShiftService(ctx context.Context, js svc.JadwalShift) error {
	_, err := jse.UpdateJadwalShiftEndpoint(ctx, js)
	if err != nil {
		f.Println("error pada endpoint:", err)
	}
	return err
}
