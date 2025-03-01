package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Province struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wilayahs")
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database tidak bisa diakses:", err)
	}

	r := gin.Default()
	r.GET("/provinces", GetProvinces)
	r.GET("/fetch-provinces", FetchAndStoreProvinces)

	fmt.Println("🚀 Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}

func FetchAndStoreProvinces(c *gin.Context) {
	resp, err := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		log.Println("Gagal mengambil data dari API:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Gagal membaca response API:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca response API"})
		return
	}

	var provinces []Province
	if err := json.Unmarshal(body, &provinces); err != nil {
		log.Println("Gagal decode JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal decode JSON"})
		return
	}

	if len(provinces) == 0 {
		log.Println("Data provinsi dari API kosong!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data provinsi dari API kosong"})
		return
	}

	_, err = db.Exec("TRUNCATE TABLE provinces")
	if err != nil {
		log.Println("Gagal menghapus data lama:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data lama"})
		return
	}

	for _, p := range provinces {
		_, err := db.Exec("INSERT INTO provinces (code, name) VALUES (?, ?)", p.Code, p.Name)
		if err != nil {
			log.Println("Gagal insert data:", err)
		}
	}

	fmt.Println("Data provinsi berhasil diperbarui di database!")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "Successfully fetched and updated data",
	})
}

func GetProvinces(c *gin.Context) {
	rows, err := db.Query("SELECT id, code, name FROM provinces")
	if err != nil {
		log.Println("Gagal mengambil data dari database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var provinces []Province
	for rows.Next() {
		var p Province
		if err := rows.Scan(&p.ID, &p.Code, &p.Name); err != nil {
			log.Println("Gagal membaca data:", err)
			continue
		}
		provinces = append(provinces, p)
	}

	response := struct {
		Status  string     `json:"status"`
		Code    int        `json:"code"`
		Message string     `json:"message"`
		Data    []Province `json:"data"`
	}{
		Status:  "success",
		Code:    200,
		Message: "Successfully get data",
		Data:    provinces,
	}

	c.JSON(http.StatusOK, response)
}
