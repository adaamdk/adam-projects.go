package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 2.
const (
	addUser        = `insert into tbUser (idUser, username, password, idKaryawan, status, createdBy, createdOn, updatedBy, updatedOn, keterangan)values (?,?,?,?,?,?,?,?,?,?)`
	updateUser     = `update tbUser set username=?, password=?, idKaryawan=?, status=?, createdBy=?, createdOn=?, updatedBy=?, updatedOn=?, keterangan=? where idUser=?`
	selectUser     = `select idUser, username, password, idKaryawan, status, createdBy, createdOn, updatedBy, updatedOn, keterangan from tbUser`
	selectUserByID = `select idUser, username, password, idKaryawan, status, createdBy, createdOn, updatedBy, updatedOn, keterangan from tbUser where idUser=?`
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

func (rw *dbReadWriter) AddUser(u User) error {
	//fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(addUser, u.ID, u.Username, u.Pwd, u.IDkaryawan, OnAdd, u.CreatedBy, time.Now(), u.UpdatedBy, time.Now(), u.Keterangan)

	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) UpdateUser(u User) error {
	fmt.Println("Update Anda berhasil gan! Silahkan lihat DB Anda.")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateUser, u.Username, u.Pwd, u.IDkaryawan, OnAdd, u.CreatedBy, time.Now(), u.UpdatedBy, time.Now(), u.Keterangan, u.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}
func (rw *dbReadWriter) ReadUser() (Users, error) {
	user := Users{}
	rows, _ := rw.db.Query(selectUser)
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Pwd, &u.IDkaryawan, &u.Status, &u.CreatedBy, &u.CreatedOn, &u.UpdatedBy, &u.UpdatedOn, &u.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return user, err
		}
		user = append(user, u)
	}
	//fmt.Println("db nya:", User)
	return user, nil
}

func (rw *dbReadWriter) ReadUserByID(ID string) (User, error) {
	fmt.Println("show by user")
	user := User{ID: ID}
	err := rw.db.QueryRow(selectUserByID, ID).Scan(&user.ID, &user.Username, &user.Pwd, &user.IDkaryawan,
		&user.Status, &user.CreatedBy, &user.CreatedOn, &user.UpdatedBy, &user.UpdatedOn, &user.Keterangan)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
