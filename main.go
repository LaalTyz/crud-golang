package main

import (
	"log"

	"github.com/LaalTyz/crud-golang/pkg/db"
	"github.com/LaalTyz/crud-golang/pkg/handlers"
)

func main() {
	// Membuat koneksi dengan database
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inisialisasi handler produk
	productHandler := handlers.NewProductHandler(db)

	// Jalankan server HTTP
	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(productHandler.RunServer())
}
