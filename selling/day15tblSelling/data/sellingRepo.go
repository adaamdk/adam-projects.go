package data

import (
	"database/sql"
	"day15tblSelling/models"
	f "fmt"
)

type SellingRepo struct {
	DB *sql.DB
}

func GetAll(db *SellingRepo) []models.Selling {
	result, err := db.DB.Query("SELECT Invoice, InvoiceDate FROM tblSelling")

	if err != nil {
		return nil
	}

	defer result.Close()
	f.Println(result)
	selling := []models.Selling{}

	for result.Next() {
		var s models.Selling

		if err := result.Scan(&s.Invoice, &s.InvoiceDate); err != nil {
			return nil
		}

		selling = append(selling, s)
	}
	return selling
}
