package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Province struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Response struct {
	Status  string     `json:"status"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Province `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/provinces", FetchProvinces)

	fmt.Println("🚀 Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}

func FetchProvinces(c *gin.Context) {
	resp, err := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		log.Println("❌ Gagal mengambil data dari API:", err)
		c.IndentedJSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Code:    500,
			Message: "Gagal mengambil data dari API",
			Data:    nil,
		})
		return
	}
	defer resp.Body.Close()

	var rawProvinces []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&rawProvinces); err != nil {
		log.Println("❌ Gagal decode JSON:", err)
		c.IndentedJSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Code:    500,
			Message: "Gagal decode JSON",
			Data:    nil,
		})
		return
	}

	var provinces []Province
	for i, p := range rawProvinces {
		provinces = append(provinces, Province{
			ID:   i + 1,
			Code: p.ID,
			Name: p.Name,
		})
	}

	c.IndentedJSON(http.StatusOK, Response{
		Status:  "success",
		Code:    200,
		Message: "Successfully get data",
		Data:    provinces,
	})
}
