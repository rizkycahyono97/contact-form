
# 🗂️ CMS Contact Form

Laravel-based CMS untuk melihat dan mengelola data contact form yang dikirim oleh pengguna dari frontend.

---

## Overview
![Desktop](storage/app/public/Screenshot%20from%202025-05-07%2008-48-15.png)

## 📦 Fitur

- Autentikasi admin
- Tabel kontak dengan pagination
- Pencarian dan filtering data
- Terintegrasi dengan REST API `api-contact-form`
- Docker-ready setup

---

## ⚙️ Requirements

- PHP 8.x
- Composer
- Laravel 10.x
- Docker & Docker Compose (opsional)

---

## 🚀 Instalasi (Local Manual)

1. **Clone project**
```bash
git clone https://github.com/yourusername/cms-contact-form.git
cd cms-contact-form
````

2. **Install dependencies**

```bash
composer install
```

3. **Copy `.env` dan generate key**

```bash
cp .env.example .env
php artisan key:generate
```

4. **Atur koneksi database di `.env`**

```env
SESSION_DRIVER=file
```

5. **Jalankan server**

```bash
php artisan serve
```

Akses: [http://127.0.0.1:8000](http://127.0.0.1:8000)

---

## 🐳 Instalasi via Docker

1. **Build dan jalankan container**

```bash
docker compose build
docker compose up -d
```

4. **Akses CMS**

* URL: [http://localhost:8081](http://localhost:8081)

---


## 🛠️ Pengaturan Tambahan

* File konfigurasi PHP lokal: `./docker/php/local.ini`
* Nginx config: `./docker/nginx/nginx.conf`
* Base API yang digunakan bisa diatur di `/config/services` dengan:

```env
'contacts_api' => [
        'base_url' => env('API_CONTACT_FORM_BASE_URI', 'http://localhost:8080/contacts'),
    ]
```

---

## 📄 License

MIT License

---

## ✍️ Author

Rizky Cahyono Putra