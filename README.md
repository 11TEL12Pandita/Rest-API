# Rest API untuk Data Provinsi Indonesia

Proyek ini adalah REST API sederhana yang mengambil data provinsi dari sumber eksternal, menyimpannya ke database MySQL, dan menyediakan endpoint untuk mengakses data tersebut.

## Fitur
- Mengambil data provinsi dari API eksternal.
- Menyimpan data ke dalam database MySQL.
- Menampilkan daftar provinsi yang tersimpan dalam database.

## Teknologi yang Digunakan
- Golang
- Gin (framework untuk web)
- MySQL
- GORM (ORM untuk database)

## Persiapan Sebelum Menjalankan
1. **Install Golang**: Pastikan Golang sudah terinstall di komputer kamu.
2. **Setup MySQL**: Buat database baru dengan nama `wilayahs`.
3. **Buat tabel `provinces` di MySQL**:
    ```sql
    CREATE TABLE provinces (
        id INT AUTO_INCREMENT PRIMARY KEY,
        code VARCHAR(10) NOT NULL,
        name VARCHAR(100) NOT NULL
    );
    ```
4. **Install package yang diperlukan**:
    ```sh
    go mod init Rest-API
    go get github.com/gin-gonic/gin
    go get github.com/go-sql-driver/mysql
    ```

## Cara Menjalankan
1. **Clone repository ini**:
    ```sh
    git clone https://github.com/11TEL12Pandita/Rest-API.git
    cd Rest-API
    ```
2. **Jalankan aplikasi**:
    ```sh
    go run main.go
    ```
3. **API akan berjalan di `http://localhost:8080`**

## Endpoint API
### 1. Fetch Data Provinsi dari API Eksternal
- **Endpoint**: `GET /fetch-provinces`
- **Deskripsi**: Mengambil data provinsi dari API eksternal dan menyimpannya ke dalam database.
- **Response**:
    ```json
    {
      "status": "success",
      "code": 200,
      "message": "Successfully fetched and updated data"
    }
    ```

### 2. Mendapatkan Daftar Provinsi dari Database
- **Endpoint**: `GET /provinces`
- **Deskripsi**: Mengambil data provinsi dari database.
- **Response**:
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
        }
      ]
    }
    ```

## Penjelasan Singkat Kode
- **`main.go`**:
  - Menghubungkan ke database.
  - Menjalankan server menggunakan Gin.
  - Menyediakan endpoint `/provinces` dan `/fetch-provinces`.
- **`FetchAndStoreProvinces()`**:
  - Mengambil data provinsi dari API eksternal.
  - Menyimpan data ke database.
- **`GetProvinces()`**:
  - Mengambil data provinsi dari database dan mengirimkannya dalam format JSON.

## Selesai 🎉
Sekarang API kamu sudah bisa digunakan! Kalau ada pertanyaan atau kendala, tinggal tanya aja. 😃

