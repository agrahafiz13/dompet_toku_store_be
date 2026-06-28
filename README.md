# 🚀 Dompet Toku - E-Money Backend dengan 2FA

Dompet Toku merupakan aplikasi backend E-Money (dompet digital) berbasis **Golang** dengan REST API yang mendukung **Two-Factor Authentication (2FA)** menggunakan Firebase OTP, Email OTP, dan TOTP (Google Authenticator).

---

# 📂 Repository

## Backend Repository

https://github.com/agrahafiz13/dompet_toku_be

---

# 🌐 Backend API

Backend dibangun menggunakan **Golang**, **Gin Framework**, **GORM**, serta **Firebase Admin SDK**, **Redis**, dan **MySQL**.

## Fitur Utama Backend:

- 🔐 **Firebase Authentication** - Login dengan Firebase
- 🔑 **Two-Factor Authentication (2FA)**
  - 📱 Firebase OTP via Push Notification
  - 📧 Email OTP
  - 🔐 TOTP (Google Authenticator / Time-based One-Time Password)
- 💳 **Manajemen Account** - Saldo digital pengguna
- 💰 **Transfer Money** - Transfer uang antar pengguna dengan OTP verification
- 💸 **Top Up** - Penambahan saldo (untuk testing)
- 📊 **Riwayat Transaksi** - Melihat history transfer dan top up
- 🔒 **JWT Authentication** - Session management dengan JWT
- 📧 **Email Service** - Pengiriman OTP via email
- 🚀 **RESTful API** - API yang clean dan terstruktur
- 🛡️ **Security Middleware** - JWT validation dan request logging

---

# 🏗️ Backend Architecture

```
backend/
│
├── config/
│   └── config.go              # Konfigurasi aplikasi
│
├── database/
│   ├── firebase.go            # Inisialisasi Firebase
│   ├── mysql.go               # Koneksi MySQL + Auto-migrate
│   └── redis.go               # Koneksi Redis
│
├── handlers/
│   ├── auth.go                # Auth & Token verification
│   ├── health.go              # Health check endpoint
│   ├── otp.go                 # OTP operations (Firebase, Email, TOTP)
│   └── payment.go             # Account & Transfer operations
│
├── middleware/
│   ├── jwt.go                 # JWT authentication middleware
│   └── logger.go              # Request/Response logging
│
├── models/
│   ├── user.go                # User model
│   ├── otp.go                 # OTP model
│   └── account.go             # Account/Balance model
│
├── services/
│   ├── email.go               # Email sending service
│   ├── firebase_rest.go       # Firebase REST API integration
│   ├── jwt.go                 # JWT generation & validation
│   └── otp.go                 # OTP generation & verification
│
├── routes/
│   └── routes.go              # API route registration
│
├── .env                       # Environment variables
├── firebase_service_account.json # Firebase credentials
├── go.mod                     # Go module definition
├── go.sum                     # Go dependency lock
└── main.go                    # Entry point
```

---

# 📁 Penjelasan Struktur Backend

## config/

Folder konfigurasi aplikasi.

| File | Fungsi |
|------|---------|
| config.go | Membaca environment variables dan menyimpan konfigurasi aplikasi |

**Konfigurasi yang disimpan:**
- Server Port
- Database Connection (MySQL)
- Cache (Redis)
- JWT Secret & Expiry
- Firebase Credentials
- SMTP Configuration
- OTP Expiry Time

---

## database/

Layer inisialisasi database dan external services.

| File | Fungsi |
|------|---------|
| firebase.go | Inisialisasi Firebase Admin SDK untuk authentication & messaging |
| mysql.go | Setup koneksi MySQL dengan GORM dan auto-migration |
| redis.go | Setup Redis client untuk caching OTP codes |

---

## handlers/

Layer yang menerima HTTP Request dari Client.

| File | Fungsi |
|------|---------|
| auth.go | Login, token verification, registration, email OTP verification |
| health.go | Health check endpoint untuk monitoring |
| otp.go | Send Firebase OTP, Send Email OTP, Confirm OTP, Register/Verify TOTP |
| payment.go | Get account, get transactions, top up, transfer |

Handler bertugas:
- Menerima HTTP request
- Validasi input
- Memanggil service business logic
- Mengembalikan JSON response

---

## middleware/

Berisi middleware aplikasi yang dijalankan sebelum handler.

| File | Fungsi |
|------|---------|
| jwt.go | Memverifikasi JWT token di Authorization header sebelum request diproses |
| logger.go | Logging request/response untuk debugging dan monitoring |

---

## models/

Representasi tabel database menggunakan GORM.

| File | Fungsi |
|------|---------|
| user.go | Model User dengan Firebase UID, email, password hash |
| otp.go | Model OTP dengan expiry time dan verification status |
| account.go | Model Account/Balance untuk setiap user |

---

## services/

Berisi seluruh Business Logic aplikasi.

| File | Fungsi |
|------|---------|
| jwt.go | JWT generation & validation dengan custom claims |
| otp.go | OTP generation, storage di Redis, verification, TOTP handling |
| email.go | Email sending menggunakan SMTP (Gmail) |
| firebase_rest.go | Firebase REST API untuk email verification link |

Service menjadi penghubung antara Handler dan Database.

---

## routes/

Mengatur seluruh endpoint API dan middleware.

| File | Fungsi |
|------|---------|
| routes.go | Registrasi seluruh endpoint beserta middleware |

---

## main.go

Merupakan entry point aplikasi.

Fungsi:
- Load konfigurasi dari .env
- Koneksi Database (MySQL)
- Inisialisasi Cache (Redis)
- Inisialisasi Firebase
- Registrasi Router
- Menjalankan HTTP Server

---

# 🔄 Request Flow

```
Flutter/Client App
        │
        ▼
HTTP Request + JWT Token
        │
        ▼
Router (Gin)
        │
        ▼
Middleware (Logger, JWT Auth)
        │
        ▼
Handler
        │
        ▼
Service (Business Logic)
        │
        ▼
Repository/Database (GORM)
        │
        ▼
MySQL Database
        │
        ▼
JSON Response
```

---

# 🔄 Alur Kerja Aplikasi

## 1. Registration & Login

```
User
  │
  ▼
POST /v1/auth/register (Firebase Token)
  │
  ▼
Backend Verify Firebase Token
  │
  ▼
Create User in Database
  │
  ▼
Generate JWT Token
  │
  ▼
Send OTP to Email
  │
  ▼
JWT Token returned
```

---

## 2. Email Verification dengan OTP

```
User menerima OTP via Email
  │
  ▼
User input OTP code
  │
  ▼
POST /v1/auth/verify-email-otp (JWT required)
  │
  ▼
Backend verify OTP dari Redis
  │
  ▼
Update user email_verified = true
  │
  ▼
Sync ke Firebase
  │
  ▼
Success Response
```

---

## 3. Transfer dengan 2FA

```
User di dashboard
  │
  ▼
Request OTP (Firebase, Email, atau TOTP)
  │
  ▼
POST /v1/otp/send-firebase | send-email
  │
  ▼
Backend send OTP via Firebase Messaging atau Email
  │
  ▼
OTP stored in Redis (5 minutes expiry)
  │
  ▼
User menerima OTP
  │
  ▼
User submit transfer + OTP code
  │
  ▼
POST /v1/payment/transfer
  │
  ▼
Backend verify OTP
  │
  ▼
Verify balance cukup
  │
  ▼
Deduct saldo pengirim
  │
  ▼
Add saldo penerima
  │
  ▼
Create transaction record
  │
  ▼
Return success response
```

---

## 4. TOTP Setup & Login dengan Google Authenticator

```
User enable TOTP
  │
  ▼
POST /v1/otp/totp/register
  │
  ▼
Backend generate TOTP secret
  │
  ▼
Generate QR code
  │
  ▼
Return secret + QR code to user
  │
  ▼
User scan QR code di Google Authenticator
  │
  ▼
POST /v1/otp/totp/verify (code from GA)
  │
  ▼
Backend verify code
  │
  ▼
Enable TOTP for user
  │
  ▼
Success - TOTP activated
```

---

# 🛠️ Tech Stack

## Backend

- **Golang** - Programming Language
- **Gin Framework** - Web framework
- **GORM** - ORM untuk database
- **MySQL** - Primary database
- **Redis** - Cache & OTP storage
- **Firebase Admin SDK** - Authentication & Cloud Messaging
- **JWT** - Session management
- **Gomail** - Email sending

## Tools

- Git
- Postman
- VS Code
- Docker (optional)

---

# 🚀 Cara Setup & Menjalankan Project

## Prerequisites

Pastikan sudah install:
- **Go 1.21+** - https://golang.org/doc/install
- **MySQL 8.0+** - https://dev.mysql.com/downloads/mysql/
- **Redis** - https://redis.io/download
- **Git** - https://git-scm.com/download

---

## 1. Clone Repository

```bash
git clone https://github.com/agrahafiz13/dompet_toku_be.git
cd dompet_toku_be
```

---

## 2. Setup Environment Variables

Copy `.env.example` menjadi `.env` dan sesuaikan konfigurasi:

```bash
cp .env.example .env
```

Edit file `.env`:

```env
# Server
PORT=8081

# Database MySQL
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=emoney_2fa

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this
JWT_EXPIRY_HOURS=24

# Firebase
FIREBASE_CREDENTIALS_PATH=firebase_service_account.json
FIREBASE_API_KEY=your_firebase_api_key

# SMTP Email
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASSWORD=your_app_password
SMTP_FROM=your_email@gmail.com
SMTP_FROM_NAME=E-Money App

# OTP
OTP_EXPIRY_MINUTES=5
```

---

## 3. Setup Firebase

1. Buat project di [Firebase Console](https://console.firebase.google.com)
2. Download Firebase Service Account JSON
3. Simpan sebagai `firebase_service_account.json` di root project
4. Update `FIREBASE_API_KEY` di `.env`

---

## 4. Setup Database MySQL

```bash
# Create database
mysql -u root -p
```

```sql
CREATE DATABASE emoney_2fa;
USE emoney_2fa;
```

Atau gunakan tool seperti MySQL Workbench / phpMyAdmin.

---

## 5. Install Dependencies

```bash
go mod tidy
```

---

## 6. Start Redis Server

```bash
# Windows (jika sudah install Memurai atau Redis)
redis-server

# Linux
redis-server

# macOS
brew services start redis
```

---

## 7. Run Backend Server

```bash
go run main.go
```

Jika berhasil, akan muncul output:

```
Firebase initialized
MySQL connected and migrated
Redis connected
Server running on :8081
```

Backend berjalan pada: `http://localhost:8081`

---

# 📡 API Endpoint

Base URL: `http://localhost:8081/v1`

Authentication Header:
```
Authorization: Bearer <jwt_token>
```

## Health Check

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| GET | `/health` | ❌ | Health check server |

---

## Authentication

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| POST | `/auth/verify-token` | ❌ | Verify Firebase token & generate JWT |
| POST | `/auth/register` | ❌ | Register user & send OTP email |
| GET | `/auth/me` | ✅ | Get current user info |
| PUT | `/auth/fcm-token` | ✅ | Update FCM token untuk push notification |
| POST | `/auth/verify-email-otp` | ✅ | Verify email OTP |

### Example Request: Register

```bash
curl -X POST http://localhost:8081/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "firebase_token": "eyJhbGciOiJSUzI1NiIs..."
  }'
```

### Example Response

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user_id": 1,
    "firebase_uid": "abc123def456",
    "email": "user@example.com",
    "jwt_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

---

## OTP - Firebase Notification

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| POST | `/otp/send-firebase` | ✅ | Send OTP via Firebase Cloud Messaging |
| POST | `/otp/confirm` | ✅ | Verify OTP code |

### Example Request: Send Firebase OTP

```bash
curl -X POST http://localhost:8081/v1/otp/send-firebase \
  -H "Authorization: Bearer <jwt_token>" \
  -H "Content-Type: application/json"
```

### Example Response

```json
{
  "success": true,
  "message": "OTP berhasil dikirim via notifikasi Firebase",
  "data": {
    "otp_type": "firebase",
    "expires_in": 300
  }
}
```

---

## OTP - Email

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| POST | `/otp/send-email` | ✅ | Send OTP via Email |

### Example Request: Send Email OTP

```bash
curl -X POST http://localhost:8081/v1/otp/send-email \
  -H "Authorization: Bearer <jwt_token>" \
  -H "Content-Type: application/json"
```

### Example Response

```json
{
  "success": true,
  "message": "OTP berhasil dikirim ke email user@example.com",
  "data": {
    "otp_type": "email",
    "expires_in": 300
  }
}
```

---

## OTP - TOTP (Google Authenticator)

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| POST | `/otp/totp/register` | ✅ | Generate TOTP secret & QR code |
| POST | `/otp/totp/verify` | ✅ | Verify TOTP code & enable TOTP |

### Example Request: Register TOTP

```bash
curl -X POST http://localhost:8081/v1/otp/totp/register \
  -H "Authorization: Bearer <jwt_token>" \
  -H "Content-Type: application/json"
```

### Example Response

```json
{
  "success": true,
  "data": {
    "secret": "JBSWY3DPEBLW64TMMQ======",
    "qr_code_base64": "data:image/png;base64,iVBORw0KG..."
  }
}
```

---

## Account

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| GET | `/account` | ✅ | Get current user account & balance |
| GET | `/account/transactions` | ✅ | Get transaction history |

### Example Request: Get Account

```bash
curl -X GET http://localhost:8081/v1/account \
  -H "Authorization: Bearer <jwt_token>"
```

### Example Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "balance": 1000000,
    "created_at": "2024-01-15T10:30:00Z"
  }
}
```

---

## Payment - Transfer

| Method | Endpoint | Auth | Deskripsi |
|----------|----------|------|-----------|
| POST | `/payment/transfer` | ✅ | Transfer uang dengan OTP verification |
| POST | `/payment/topup` | ✅ | Top up saldo (untuk testing) |

### Example Request: Transfer

```bash
curl -X POST http://localhost:8081/v1/payment/transfer \
  -H "Authorization: Bearer <jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "recipient_account_number": "1234567890",
    "amount": 100000,
    "description": "Pembayaran tagihan",
    "otp_code": "123456",
    "otp_type": "firebase"
  }'
```

### Example Response

```json
{
  "success": true,
  "message": "Transfer berhasil",
  "data": {
    "transaction_id": "TRX123456",
    "status": "SUCCESS",
    "amount": 100000,
    "balance": 900000,
    "timestamp": "2024-01-15T10:35:00Z"
  }
}
```

---

# 🔗 Integrasi dengan Frontend

Frontend menggunakan REST API dengan format JSON.

### Base URL
```
http://localhost:8081/v1
```

### Authentication Flow

1. **Login dengan Firebase**
   ```
   POST /auth/verify-token
   ```
   - Kirim Firebase ID Token
   - Backend mengembalikan JWT Token

2. **Gunakan JWT Token untuk request selanjutnya**
   ```
   Authorization: Bearer <jwt_token>
   ```

3. **Token akan ter-decode di middleware**
   - Ambil `user_id` dari JWT claims
   - Lanjutkan request ke handler

---

# 🔐 Security Features

## JWT Token Protection
- Token di-sign dengan secret key
- Token memiliki expiry time (default 24 jam)
- Setiap request protected memerlukan valid JWT token

## OTP Security
- OTP di-store di Redis (encrypted in transit)
- OTP expire dalam 5 menit
- Max attempt verification (3x gagal = lock)

## TOTP (Time-based OTP)
- Generated menggunakan TOTP algorithm
- Diverifikasi dengan Google Authenticator
- Backup code untuk recovery

## Password Security
- Password di-hash sebelum disimpan
- Firebase authentication untuk secure login
- Email verification sebelum full access

## Request Logging
- Semua request/response di-log
- Sensitive fields di-mask (password, token, otp_code)
- Log untuk debugging & security audit

---

# 📋 Postman Collection

Gunakan Postman untuk testing API. Import file:
```
postman/emoney-2fa.postman_collection.json
```

Atau buat manual dengan endpoint yang sudah dijelaskan di atas.

---

# 🐛 Troubleshooting

### Error: "Failed to connect to MySQL"
- Pastikan MySQL service berjalan
- Cek DB_HOST, DB_PORT, DB_USER, DB_PASSWORD di .env
- Pastikan database sudah dibuat

### Error: "Failed to connect to Redis"
- Pastikan Redis server berjalan
- Cek REDIS_HOST, REDIS_PORT di .env
- Jalankan `redis-server` di terminal

### Error: "Firebase not initialized"
- Pastikan `firebase_service_account.json` ada di root project
- Cek path di FIREBASE_CREDENTIALS_PATH
- Verify Firebase credentials valid

### Error: "SMTP not configured"
- Pastikan SMTP_USER dan SMTP_PASSWORD terisi di .env
- Untuk Gmail: gunakan App Password, bukan regular password
- Enable "Less secure app access" atau generate App Password

### OTP tidak diterima via Email
- Cek SMTP configuration
- Check spam folder
- Cek SMTP_FROM email address

---

# 📚 Additional Resources

- [Gin Framework Documentation](https://gin-gonic.com/)
- [GORM Documentation](https://gorm.io/)
- [Firebase Admin SDK Go](https://firebase.google.com/docs/database/admin/start)
- [Redis Documentation](https://redis.io/documentation)
- [JWT.io](https://jwt.io/)

---

# 👨‍💻 Developer

**Nama :** Agra Alfian Hafiz

**NIM :** 1123150025

**Kelas :** TI SE 23 M

---

# 📄 Lisensi

Project ini dibuat untuk keperluan pembelajaran, penelitian, dan tugas akademik.

---

## ⭐ Catatan Penting

- **Jangan commit `.env`** - Gunakan `.env.example` dan add ke `.gitignore`
- **Jangan commit `firebase_service_account.json`** - Ini credentials sensitif
- **Ganti JWT_SECRET** dengan secret yang kuat di production
- **Ganti SMTP_PASSWORD** dengan App Password Gmail, bukan regular password
- **Jalankan MySQL & Redis** sebelum backend
- **Setup Firebase project** sebelum test authentication
- **Gunakan Postman** untuk testing API endpoint
- **Setup CORS** jika frontend di domain berbeda

---

## 🚀 Development Tips

1. **Hot Reload Development** (Optional)
   ```bash
   go install github.com/cosmtrek/air@latest
   air
   ```

2. **Database Migration**
   - Semua table auto-migrate saat startup
   - Modifikasi models di `models/` folder
   - Restart server untuk re-migrate

3. **Debugging**
   - Enable request logging di middleware
   - Check database dengan MySQL client
   - Monitor Redis key dengan `redis-cli`

4. **Testing**
   - Gunakan Postman collection yang sudah disediakan
   - Test flow dari login → OTP → Transfer
   - Verify database changes setelah setiap operasi

---

**Last Updated:** June 2024
**Version:** 1.0.0
