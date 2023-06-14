package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@localhost(localhost:3306)/crud-golang")
	if err != nil {
		return nil, fmt.Errorf("koneksi database gagal: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("koneksi database gagal: %v", err)
	}

	return db, nil
}
