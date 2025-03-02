# Rest API untuk Data Provinsi Indonesia

Proyek ini adalah REST API sederhana yang digunakan untuk mengambil data provinsi di Indonesia dari sumber eksternal dan menyimpannya ke database MySQL. API ini dibuat menggunakan bahasa pemrograman Go dengan framework **Gin**.

## ✨ Fitur
- **Ambil Data Provinsi dari API**: Mengambil data provinsi dari API eksternal.
- **Simpan Data ke Database**: Menyimpan data ke database MySQL.
- **Ambil Data dari Database**: Mengambil data provinsi yang sudah tersimpan di database.

---

## 🚀 Cara Menjalankan Proyek

### 1️⃣ Persiapan
Sebelum menjalankan proyek ini, pastikan sudah menginstal:
- Go (minimal versi 1.17)
- MySQL (pastikan sudah membuat database `wilayahs` dengan tabel `provinces`)
- Git

Kemudian, clone repository ini:
```sh
$ git clone https://github.com/11TEL12Pandita/Rest-API.git
$ cd Rest-API
```

### 2️⃣ Instalasi Dependensi
Jalankan perintah berikut untuk menginstal dependensi yang dibutuhkan:
```sh
$ go mod tidy
```

### 3️⃣ Konfigurasi Database
Pastikan database MySQL sudah berjalan, lalu buat database dan tabel dengan struktur berikut:
```sql
CREATE DATABASE wilayahs;
USE wilayahs;
CREATE TABLE provinces (
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name VARCHAR(255) NOT NULL
);
```

### 4️⃣ Menjalankan Server
Jalankan perintah berikut untuk menjalankan server:
```sh
$ go run main.go
```
Server akan berjalan di `http://localhost:8080`

---

## 🔗 Endpoint API

### 1️⃣ Ambil Semua Provinsi dari Database
- **Endpoint**: `GET /provinces`
- **Deskripsi**: Mengambil semua provinsi yang sudah tersimpan di database.
- **Response Contoh**:
  ```json
  {
    "status": "success",
    "code": 200,
    "message": "Successfully get data",
    "data": [
      {
        "id": 1,
        "code": "11",
        "name": "Aceh"
      },
      {
        "id": 2,
        "code": "12",
        "name": "Sumatera Utara"
      }
    ]
  }
  ```

### 2️⃣ Fetch Data Provinsi dari API Eksternal dan Simpan ke Database
- **Endpoint**: `GET /fetch-provinces`
- **Deskripsi**: Mengambil data dari API eksternal dan menyimpannya ke database MySQL.
- **Response Contoh**:
  ```json
  {
    "status": "success",
    "code": 200,
    "message": "Successfully fetched and updated data"
  }
  ```

---

## 🛠 Penjelasan Kode

### 📌 Struktur Kode
- **Struct `Province`**: Digunakan untuk merepresentasikan data provinsi.
- **Fungsi `main`**: Menjalankan server menggunakan framework Gin.
- **Fungsi `FetchAndStoreProvinces`**: Mengambil data dari API eksternal dan menyimpannya ke database.
- **Fungsi `GetProvinces`**: Mengambil data dari database dan mengembalikannya dalam format JSON.

### 📝 Contoh Potongan Kode
**Mengambil Data dari API dan Menyimpannya ke Database:**
```go
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
```

---

## 🤝 Kontribusi
Jika ingin berkontribusi atau memiliki masukan, silakan fork repository ini dan buat pull request! 🚀

---

## 📜 Lisensi
Proyek ini menggunakan lisensi MIT. Silakan gunakan dan modifikasi sesuai kebutuhan.

---

### 🔥 Selamat ngoding! 🚀

