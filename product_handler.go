package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LaalTyz/crud-golang/pkg/models"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	DB *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (ph *ProductHandler) RunServer() error {
	r := mux.NewRouter()

	r.HandleFunc("/products", ph.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", ph.GetProductByID).Methods("GET")
	r.HandleFunc("/products", ph.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", ph.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", ph.DeleteProduct).Methods("DELETE")

	return http.ListenAndServe(":8080", r)
}

func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan daftar produk dari database
	products, err := ph.fetchProducts()
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mengambil data produk", http.StatusInternalServerError)
		return
	}

	// Mengirimkan daftar produk sebagai response JSON
	json.NewEncoder(w).Encode(products)
}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "ID produk tidak valid", http.StatusBadRequest)
		return
	}

	// Mendapatkan produk dari database berdasarkan ID
	product, err := ph.fetchProductByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mengambil data produk", http.StatusInternalServerError)
		return
	}

	// Mengirimkan produk sebagai response JSON
	json.NewEncoder(w).Encode(product)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, "Data produk tidak valid", http.StatusBadRequest)
		return
	}

	// Membuat produk baru di database
	err = ph.createProduct(product)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal membuat produk", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produk berhasil dibuat")
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "ID produk tidak valid", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, "Data produk tidak valid", http.StatusBadRequest)
		return
	}

	product.ID = id

	// Mengupdate produk di database
	err = ph.updateProduct(product)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mengupdate produk", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Produk berhasil diperbarui")
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "ID produk tidak valid", http.StatusBadRequest)
		return
	}

	// Menghapus produk dari database
	err = ph.deleteProduct(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal menghapus produk", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Produk berhasil dihapus")
}

// Implementasi fungsi-fungsi akses database (fetch, create, update, delete)
// Sesuaikan dengan logika akses database Anda
func (ph *ProductHandler) fetchProducts() ([]models.Product, error) {
	// Implementasi logika akses database untuk mendapatkan daftar produk
}

func (ph *ProductHandler) fetchProductByID(id int) (models.Product, error) {
	// Implementasi logika akses database untuk mendapatkan produk berdasarkan ID
}

func (ph *ProductHandler) createProduct(product models.Product) error {
	// Implementasi logika akses database untuk membuat produk baru
}

func (ph *ProductHandler) updateProduct(product models.Product) error {
	// Implementasi logika akses database untuk mengupdate produk
}

func (ph *ProductHandler) deleteProduct(id int) error {
	// Implementasi logika akses database untuk menghapus produk
}
