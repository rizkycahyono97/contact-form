# Client Contact Form (Next.js)

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Installation](#installation)
5. [Configuration](#configuration)
6. [Running the Application](#running-the-application)
7. [API Endpoints](#api-endpoints)
8. [Testing](#testing)
9. [Contributing](#contributing)
10. [License](#license)

---

## Overview
![Desktop](public/Screenshot%20from%202025-05-07%2008-44-03.png)

## Introduction

Aplikasi **Client Contact Form** adalah frontend berbasis **Next.js** yang digunakan untuk mengirim data formulir kontak ke backend [**API Contact Form**](https://github.com/rizkycahyono97/contact-form/tree/main/app/api-contact-form). Aplikasi ini menyediakan antarmuka pengguna untuk mengisi formulir kontak dan mengirimkannya ke server melalui API.

---

## Features

- **Formulir Kontak**: Pengguna dapat mengisi formulir dengan detail seperti nama, email, nomor telepon, dan pesan.
- **Validasi Klien**: Memastikan semua field wajib diisi sebelum pengiriman.
- **Error Handling**: Menampilkan pesan kesalahan jika terjadi masalah validasi atau server.
- **Responsif**: Desain responsif untuk mendukung berbagai perangkat.
- **Integrasi API**: Mengirim data formulir ke backend [**API Contact Form**](https://github.com/rizkycahyono97/contact-form/tree/main/app/api-contact-form) menggunakan metode POST.

---

## Prerequisites

Sebelum menjalankan aplikasi ini, pastikan Anda telah memenuhi persyaratan berikut:

- **Node.js**: Versi LTS (18.13.0 atau lebih tinggi).
- **npm** atau **yarn**: Package manager untuk Node.js.
- **Backend API**: Pastikan [**API Contact Form**](https://github.com/rizkycahyono97/contact-form/tree/main/app/api-contact-form) sudah berjalan dan dapat diakses melalui URL tertentu (misalnya, `http://localhost:9000`).

---

## Installation

### 1. Clone Repository
Clone repository ini ke komputer lokal Anda:
```bash
git clone https://github.com/your-repo/client-contact-form.git
cd client-contact-form
```

### 2. Install Dependencies
Instal semua dependensi yang diperlukan menggunakan npm atau yarn:
```bash
npm install
# atau
yarn install
```

---

## Configuration

### 1. File `.env.local`
Buat file `.env.local` di root direktori proyek dan isi dengan variabel lingkungan berikut:
```env
NEXT_PUBLIC_API_URL=http://localhost:9000
```

- **`NEXT_PUBLIC_API_URL`**: URL backend API yang digunakan untuk mengirim data formulir.

---

## Running the Application

### 1. Start the Application in Development Mode
Jalankan aplikasi dalam mode pengembangan menggunakan perintah berikut:
```bash
npm run dev
# atau
yarn dev
```

Aplikasi akan berjalan di port `3000`. Buka browser dan navigasikan ke URL berikut:
```
http://localhost:3000
```

### 2. Build and Run in Production Mode
Untuk menjalankan aplikasi dalam mode produksi, ikuti langkah-langkah berikut:

#### Build the Application
```bash
npm run build
# atau
yarn build
```

#### Start the Application
```bash
npm start
# atau
yarn start
```

Aplikasi akan berjalan di port `3000`.

---



## Testing

### 1. Manual Testing
- Isi formulir kontak di halaman utama (`http://localhost:3000`).
- Verifikasi bahwa data berhasil dikirim ke backend API.
- Coba kirim formulir dengan input yang salah untuk memastikan validasi berfungsi.

### 2. Postman Collection
Gunakan file JSON Postman Collection untuk menguji API secara langsung. Import file koleksi dari direktori proyek:
```
app/api-contact-form/test/API Contact Form.postman_collection.json
```

---

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

## Contributing

Kontribusi sangat diterima! Jika Anda ingin berkontribusi, silakan buat pull request dengan penjelasan detail tentang perubahan yang Anda buat.

---

## License

Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk detail lebih lanjut.
