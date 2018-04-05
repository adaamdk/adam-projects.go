package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addJadwalShift    = `insert into tbShift (shiftCode, jadwalMasuk, jadwalPulang, status, createdBy, createdOn, updatedBy, updatedOn)values (?,?,?,?,?,?,?,?)`
	updateJadwalShift = `update tbShift set jadwalMasuk=?, jadwalPulang=?, status=?, createdBy=?, createdOn=?, updatedBy=?, updatedOn=? where shiftCode=?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddJadwalShift(shift JadwalShift) error {
	//fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(addJadwalShift, shift.Code, shift.JadwalOn, shift.JadwalOff, OnAdd, shift.CreatedBy, time.Now(), shift.UpdatedBy, time.Now())

	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) UpdateJadwalShift(sh JadwalShift) error {
	fmt.Println("Update Anda berhasil gan! Silahkan lihat DB Anda.")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateJadwalShift, sh.JadwalOn, sh.JadwalOff, OnAdd, sh.CreatedBy, time.Now(), sh.UpdatedBy, time.Now(), sh.Code)

	//fmt.Println("JadwalMasuk:", sh.JadwalOn, "ShiftCode:", sh.Code)
	if err != nil {
		return err
	}

	return tx.Commit()
}
