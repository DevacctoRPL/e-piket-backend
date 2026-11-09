<h1 align="center">
  <a href="https://gofiber.io">
    <picture>
      <source height="100" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
      <img height="100" alt="Fiber" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
    </picture>
  </a>
  <br>
  E-Piket Backend
</h1>

<p align="center">
  <em>Backend API untuk sistem E-Piket yang dibangun menggunakan <b>Go Fiber v3</b>, <b>GORM</b>, dan <b>MySQL</b>.</em>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/gofiber/fiber/v3"><img src="https://img.shields.io/badge/Fiber-v3.0.0-00ACD7?style=flat-square&logo=go" alt="Fiber v3"></a>
  <a href="https://gorm.io"><img src="https://img.shields.io/badge/ORM-GORM-blue?style=flat-square" alt="GORM"></a>
  <a href="https://www.mysql.com"><img src="https://img.shields.io/badge/Database-MySQL%208.0-orange?style=flat-square&logo=mysql" alt="MySQL"></a>
  <a href="https://docs.docker.com/compose/"><img src="https://img.shields.io/badge/Container-Docker%20Compose-2496ED?style=flat-square&logo=docker" alt="Docker"></a>
</p>

---

## ⚡️ Prerequisites

- **Go**: Version `1.25` atau lebih baru
- **Docker & Docker Compose** (untuk MySQL & phpMyAdmin)
- **Git**

---

## 🚀 Quickstart

### 1. Clone Repository & Install Dependensi
```bash
git clone https://github.com/your-username/e-piket-backend.git
cd e-piket-backend
go mod download
```

### 2. Setup Environment
Salin template konfigurasi `.env`:
```bash
cp .env.example .env
```

### 3. Jalankan Database (Docker Compose)
Jalankan MySQL dan phpMyAdmin di latar belakang:
```bash
docker compose up -d
```
> **phpMyAdmin**: [http://localhost:8080](http://localhost:8080)  
> Akun default: User `digitalpiket` / Password `piketsmkpenus#`

### 4. Jalankan Aplikasi
```bash
go run cmd/api/main.go
```
Server akan berjalan di `http://localhost:3000`.

---

## 📌 API Endpoints

| Method | Endpoint | Deskripsi |
|---|---|---|
| `GET` | `/` | Cek server status |
| `GET` | `/health` | Healthcheck server & database |
| `GET` | `/api/v1/users` | Mengambil seluruh data user |
| `POST` | `/api/v1/users` | Menambahkan data user baru |

### Contoh Request Tambah User:
```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Guru Piket", "email": "piket@smkpenus.sch.id", "role": "guru_piket"}'
```

---

## 🔗 Referensi

- [Dokumentasi Go Fiber v3](https://docs.gofiber.io/)
- [GORM Docs](https://gorm.io/)
