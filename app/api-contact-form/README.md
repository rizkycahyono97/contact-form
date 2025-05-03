# API Contact Form Documentation

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Installation](#installation)
5. [Configuration](#configuration)
6. [API Endpoints](#api-endpoints)
7. [Running the Application](#running-the-application)
8. [Docker Support](#docker-support)
9. [Testing](#testing)
10. [Contributing](#contributing)
11. [License](#license)

---

## Introduction

API Contact Form adalah RESTful API yang dirancang untuk mengelola data formulir kontak. Aplikasi ini memungkinkan pengguna untuk menambahkan, melihat, memperbarui, dan menghapus data kontak melalui endpoint API yang telah disediakan. Aplikasi ini menggunakan stack teknologi berbasis GOLANG dengan framework GIN dan GORM sebagai backend dan MariaDB sebagai database.

---

## Features

- **Tambah Data Kontak**: Pengguna dapat menambahkan data kontak baru.
- **Lihat Semua Data Kontak**: Pengguna dapat melihat daftar semua data kontak.
- **Lihat Detail Data Kontak**: Pengguna dapat melihat detail data kontak berdasarkan ID.
- **Perbarui Data Kontak**: Pengguna dapat memperbarui data kontak berdasarkan ID.
- **Hapus Data Kontak**: Pengguna dapat menghapus data kontak berdasarkan ID.
- **CORS Support**: API mendukung Cross-Origin Resource Sharing (CORS) untuk frontend multi-origin.
- **Persistensi Data**: Data disimpan dalam database MariaDB untuk memastikan persistensi.
- **Containerized Deployment**: Aplikasi dapat dijalankan menggunakan Docker untuk deployment mudah dan konsisten.

---

## Prerequisites

Sebelum menjalankan aplikasi ini, pastikan Anda telah memenuhi persyaratan berikut:

- **GOLANG**: Versi 1.24 / atau gunakan Docker.
- **MariaDB**: Database MariaDB harus tersedia (atau gunakan Docker untuk menjalankan MariaDB).
- **Docker** (Opsional): Untuk menjalankan aplikasi dalam container Docker.

---

## Installation

### 1. Clone Repository
Clone repository ini ke komputer lokal Anda:
```bash
git clone https://github.com/your-repo/api-contact-form.git
cd api-contact-form
```
---

## Configuration

### 1. File `.env`
Buat file `.env` di root direktori proyek dan isi dengan variabel lingkungan berikut:
```env
HOST_API_PORT=9000
CONT_API_PORT=3000
HOST_PHPMYADMIN_PORT=8080
CONT_PHPMYADMIN_PORT=80
HOST_MARIADB_PORT=3306
CONT_MARIADB_PORT=3306
MYSQL_ROOT_PASSWORD=rootpassword
MYSQL_DATABASE=contact_form_db
MYSQL_USER=contact_user
MYSQL_PASSWORD=userpassword
```

### 2. Konfigurasi CORS
Ubah variabel `CORS_ALLOWED_ORIGINS` di file `.env` sesuai dengan origin frontend Anda atau biarkan default seperti pada `.env.example`:
```env
CORS_ALLOWED_ORIGINS=http://localhost:8081,http://localhost:8082
```

---

## API Endpoints

### 1. APICheck
- **URL**: `/`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "code": "SUCCESS",
    "message": "API Contact is Runnning",
    "data": null
  }
  ```

### 2. HealthCheck
- **URL**: `/health`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "code": "Success",
    "message": "",
    "data": "API is working"
  }
  ```
  
### 3. Tambah Data Kontak
- **URL**: `/contacts`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "1234567890",
    "message": "Hello, this is a test message."
  }
  ```
  - **Response**:
      - Success (201):
        ```json
        {
        "code": "CREATED",
        "message": "Contact Created Successfully",
        "data": {
        "id": 4,
          "name": "John Doe 6",
          "email": "john@example.com",
          "phone": "1234567890",
          "message": "Hello, World!",
          "created_at": "2025-05-03 10:04:23",
          "updated_at": "2025-05-03 10:04:23"
          }
        }
        ```

### 2. Lihat Semua Data Kontak
- **URL**: `/contacts`
- **Method**: `GET`
- **Response**:
  ```json
    {
    "code": "SUCCESS",
    "message": "Contacts Retrivied Successfully",
    "data": [
        {
            "id": 2,
            "name": "John Doe 6",
            "email": "john@example.com",
            "phone": "1234567890",
            "message": "Hello, World!",
            "created_at": "2025-05-03 10:04:19",
            "updated_at": "2025-05-03 10:04:19"
        },
        {
            "id": 3,
            "name": "Jane Doe New",
            "email": "jane@example.com",
            "phone": "0987654321",
            "message": "Updated Message",
            "created_at": "2025-05-03 10:04:22",
            "updated_at": "2025-05-03 10:04:47"
        }
    ]
  }
  ```

### 3. Lihat Detail Data Kontak
- **URL**: `/contacts/:id`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "code": "SUCCESS",
    "message": "Contact Retrivied Successfully",
    "data": {
        "id": 3,
        "name": "John Doe 6",
        "email": "john@example.com",
        "phone": "1234567890",
        "message": "Hello, World!",
        "created_at": "2025-05-03 10:04:22",
        "updated_at": "2025-05-03 10:04:22"
    }
  }
  ```

### 4. Perbarui Data Kontak
- **URL**: `/contacts/:id`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
  "name": "Jane Doe New",
  "email": "jane@example.com",
  "phone": "0987654321",
  "message": "Updated Message"
  }
  ```
- **Response**:
  ```json
  {
    "code": "SUCCESS",
    "message": "Contact Updated Successfully",
    "data": {
        "id": 3,
        "name": "Jane Doe New",
        "email": "jane@example.com",
        "phone": "0987654321",
        "message": "Updated Message",
        "created_at": "2025-05-03 10:04:22",
        "updated_at": "2025-05-03 10:04:47"
    }
  }
  ```

### 5. Hapus Data Kontak
- **URL**: `/contacts/:id`
- **Method**: `DELETE`
- **Response**:
  ```json
  {
    "code": "SUCCESS",
    "message": "Contact Deleted Successfully",
    "data": null
  }
  ```

---

## Running the Application

### 1. Start the Application
Jalankan aplikasi menggunakan perintah berikut:
```bash
go build
./contact-form-api
```

## Docker Support

### 1. Build and Run with Docker
Jika Anda ingin menjalankan aplikasi menggunakan Docker, ikuti langkah-langkah berikut:

#### Build and Start Containers
```bash
docker compose build
docker compose up -d
```

#### Stop Containers
```bash
docker compose down
```

### 2. Verify Running Containers
Periksa status container:
```bash
docker ps
```

---

**## Testing

Untuk menguji API Contact Form, Anda dapat menggunakan file koleksi Postman yang telah disediakan, import di Postman anda.

```bash
test/API Contact Form.postman_collection.json
```

## Contributing

Kontribusi sangat diterima! Jika Anda ingin berkontribusi, silakan buat pull request dengan penjelasan detail tentang perubahan yang Anda buat.

---

## License

Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk detail lebih lanjut.

---