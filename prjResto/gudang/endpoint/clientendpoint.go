package endpoint

import (
	"context"
	f "fmt"

	svc "prjResto/gudang/server"
)

func (ge GudangEndpoint) AddGudangService(ctx context.Context, usr svc.Gudang) error {
	_, err := ge.AddGudangEndpoint(ctx, usr)
	return err
}

func (ge GudangEndpoint) UpdateGudangService(ctx context.Context, usr svc.Gudang) error {
	_, err := ge.UpdateGudangEndpoint(ctx, usr)
	if err != nil {
		f.Println("error pada endpoint:", err)
	}
	return err
}

func (ge GudangEndpoint) ReadGudangByKeteranganService(ctx context.Context, k string) (svc.Gudangs, error) {
	req := svc.Gudang{Keterangan: k}
	response, err := ge.ReadGudangByKeteranganEndpoint(ctx, req)
	f.Println("Response", response)
	if err != nil {
		f.Println("Error di endpoint lu gan..", err)
	}
	return response.(svc.Gudangs), err
}
