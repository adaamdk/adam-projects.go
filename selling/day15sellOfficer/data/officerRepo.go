package data

import (
	"database/sql"
	"day15sellOfficer/models"
	f "fmt"
)

type OfficerRepo struct {
	DB *sql.DB
}

func GetAll(db *OfficerRepo) []models.Officer {
	result, err := db.DB.Query("SELECT OfficerCode, OfficerName FROM tblOfficer")

	if err != nil {
		return nil
	}

	defer result.Close()
	f.Println(result)
	officer := []models.Officer{}

	for result.Next() {
		var off models.Officer

		if err := result.Scan(&off.Kode, &off.Nama); err != nil {
			return nil
		}
		officer = append(officer, off)
	}
	return officer
}
