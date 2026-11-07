# E-Piket Backend (Golang Fiber v3 + MySQL + phpMyAdmin)

Backend API untuk sistem E-Piket menggunakan framework **Go Fiber v3**, **GORM**, **MySQL 8.0**, dan **phpMyAdmin** dengan orchestrasi **Docker Compose**.

---

## 📋 Kredensial Database & phpMyAdmin

| Parameter | Nilai |
|---|---|
| **Host** | `127.0.0.1` (atau `mysql` jika dalam container docker) |
| **Port MySQL** | `3306` |
| **Database** | `digitalpiket` |
| **Username** | `digitalpiket` (atau `root`) |
| **Password** | `piketsmkpenus#` |
| **URL phpMyAdmin** | [http://localhost:8080](http://localhost:8080) |
| **URL Backend API** | [http://localhost:3000](http://localhost:3000) |

---

## 🚀 Cara Menjalankan

### 1. Jalankan MySQL & phpMyAdmin (Docker Compose)
Jalankan perintah berikut di root folder project:

```bash
docker compose up -d
```

Periksa status container:
```bash
docker compose ps
```

Container `mysql` dan `phpmyadmin` akan berjalan di background. phpMyAdmin dapat diakses di browser:
👉 **[http://localhost:8080](http://localhost:8080)**  
- Server: `mysql`
- Username: `digitalpiket` (atau `root`)
- Password: `piketsmkpenus#`

---

### 2. Jalankan Backend Go Fiber v3 (Lokal)
Jalankan aplikasi backend menggunakan Go:

```bash
go run cmd/api/main.go
```

Aplikasi akan:
1. Terhubung otomatis ke MySQL.
2. Melakukan **Auto-Migrate** tabel (contoh: tabel `users`).
3. Berjalan di `http://localhost:3000`.

---

### 3. (Opsional) Menjalankan Semua Layanan via Docker
Jika ingin menjalankan backend di dalam container juga:

```bash
docker compose --profile full up -d --build
```

---

## 📌 Endpoint API

| Method | Endpoint | Keterangan |
|---|---|---|
| `GET` | `/` | Cek status server |
| `GET` | `/health` | Healthcheck server & koneksi database |
| `GET` | `/api/v1/users` | Mengambil seluruh data user |
| `POST` | `/api/v1/users` | Menambahkan user baru |

### Contoh Request Tambah User (POST `/api/v1/users`):
```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Guru Piket", "email": "piket@smkpenus.sch.id", "role": "guru_piket"}'
```

---

## 📁 Struktur Direktori
```text
e-piket-be/
├── cmd/
│   └── api/
│       └── main.go           # Entry point aplikasi
├── config/
│   ├── config.go             # Loader konfigurasi environment (.env)
│   └── database.go           # Koneksi GORM MySQL & Auto-migrate
├── internal/
│   ├── handlers/
│   │   ├── health_handler.go # Handler health check & database ping
│   │   └── user_handler.go   # Handler CRUD user
│   ├── models/
│   │   └── user.go           # Model User GORM
│   └── routes/
│       └── routes.go         # Registrasi router & middleware
├── .env                      # File konfigurasi environment lokal
├── .env.example              # Template file konfigurasi environment
├── .gitignore
├── Dockerfile                # Dockerfile multi-stage production-ready
├── docker-compose.yml        # Docker Compose (MySQL, phpMyAdmin, Backend)
├── go.mod
└── go.sum
```
