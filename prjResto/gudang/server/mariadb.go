package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 2.
const (
	addGudang                = `insert into tbGudang (idGudang, nama, alamat, luas, status, createdBy, createdOn, updatedBy, updatedOn)values (?,?,?,?,?,?,?,?,?)`
	updateGudang             = `update tbGudang set nama=?, alamat=?, luas=?, status=?, createdBy=?, createdOn=?, updatedBy=?, updatedOn=? where idGudang=?`
	selectGudangByKeterangan = `select idGudang, nama, keterangan from tbGudang where keterangan like ?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, gudang string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", gudang, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddGudang(g Gudang) error {
	//fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(addGudang, g.ID, g.Name, g.Alamat, g.Luas, OnAdd, g.CreatedBy, time.Now(), g.UpdatedBy, time.Now())

	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) UpdateGudang(g Gudang) error {
	fmt.Println("Update Anda berhasil gan! Silahkan lihat DB Anda.")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateGudang, g.Name, g.Alamat, g.Luas, OnAdd, g.CreatedBy, time.Now(), g.UpdatedBy, time.Now(), g.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadGudangByKeterangan(k string) (Gudangs, error) {
	gudang := Gudangs{} //tambah 's'
	rows, _ := rw.db.Query(selectGudangByKeterangan, k)
	defer rows.Close()
	for rows.Next() {
		var gd Gudang
		err := rows.Scan(&gd.ID, &gd.Name, &gd.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return gudang, err
		}
		gudang = append(gudang, gd)
	}
	fmt.Println("db nya:", gudang)
	return gudang, nil
}
