# 📦 Contact Form Mini Project

Mini project berbasis **microservices** yang terdiri dari tiga bagian utama:

1. **API Contact Form** - Backend REST API untuk menangani data contact form
2. **CMS Contact Form** - CMS berbasis Laravel untuk melihat/mengelola submission
3. **Client Contact Form** - Frontend form untuk user melakukan pengisian contact

---

## 🧭 Project Structure

```

.
├── api-contact-form/      # REST API (Golang)
├── cms-contact-form/      # CMS (Laravel 12)
├── client-contact-form/   # Frontend client (Express.js)
└── docker-compose.yaml    # Global orchestration file

````
---

## 🧱 Individual Service Setup

### 📡 1. `api-contact-form`

#### Deskripsi:

REST API untuk menerima dan menyimpan contact form submission.

#### Jalankan secara mandiri:

```bash
cd api-contact-form
docker compose build
docker compose up -d 
```

---

### 🖥️ 2. `cms-contact-form` (Laravel CMS)

#### Deskripsi:

Admin panel berbasis Laravel untuk melihat data submission form.

#### Jalankan secara mandiri:

```bash
cd cms-contact-form
docker compose build
docker compose up -d
```

#### Laravel Setup:

```bash
docker compose build
docker compose up -d
exit
```

---

### 🌐 3. `client-contact-form`

#### Deskripsi:

Formulir frontend untuk user mengirim pesan ke API.

#### Jalankan secara mandiri:

```bash
cd client-contact-form
docker compose build
docker compose up -d
```

---

## 🌐 URLs

| Service       | URL                                            | Description                 |
| ------------- | ---------------------------------------------- | --------------------------- |
| Client Form   | [http://localhost:3000](http://localhost:3000) | User-facing contact form    |
| CMS Dashboard | [http://localhost:8081](http://localhost:8081) | Admin panel (Laravel-based) |
| Contact API   | [http://localhost:8080](http://localhost:8080) | REST API endpoint           |

---

## 🧪 Testing

**Test API**

```bash
curl -X POST http://localhost:8080/contacts \
     -H "Content-Type: application/json" \
     -d '{"name":"Jane", "email":"jane@example.com", "message":"Hello!"}'
```

---

## 📄 License

MIT License

---

## ✍️ Author

Aditya Firman Nugroho

```