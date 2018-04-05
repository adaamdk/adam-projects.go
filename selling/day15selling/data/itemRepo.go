package data

import (
	"database/sql"
	"day15selling/models"
	f "fmt"
	// _ "github.com/go-sql-driver/mysql"
)

type ItemRepo struct {
	DB *sql.DB
}

func GetAll(db *ItemRepo) []models.Item {

	result, err := db.DB.Query("SELECT ItemCode, ItemName FROM tblItem")

	if err != nil {
		return nil
	}

	defer result.Close()
	f.Println(result)
	item := []models.Item{}

	for result.Next() {
		var i models.Item

		if err := result.Scan(&i.Kode, &i.NamaBarang); err != nil {
			return nil
		}
		item = append(item, i)
	}
	return item
}
