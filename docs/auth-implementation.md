# Authentication & JWT Implementation Guide

## Summary

✅ **Proses Auth JWT telah selesai diimplementasi dengan fitur:**

### 1. **Login Endpoint** 
- **POST** `/api/v1/login`
- Input: email & password
- Output: access_token & refresh_token

### 2. **Refresh Token Endpoint**
- **POST** `/api/v1/refresh-token` 
- Input: refresh_token
- Output: access_token baru

### 3. **Protected Endpoints**
- Semua endpoint user (`GET`, `PUT`, `DELETE /api/v1/users`) dilindungi JWT
- Harus mengirim header: `Authorization: Bearer <access_token>`
- **🔐 Swagger sudah menampilkan field Authorization untuk endpoint protected**

### 4. **Public Endpoints**
- `POST /api/v1/login` - Login
- `POST /api/v1/refresh-token` - Refresh token
- `POST /api/v1/users` - Register user (create user)

---

## 🎯 **Swagger Authentication**

### **Cara menggunakan token di Swagger:**

1. **Klik tombol "Authorize" 🔒** di kanan atas Swagger UI
2. **Masukkan token** dalam format: `Bearer <your_access_token>`
3. **Klik "Authorize"**
4. **Sekarang semua endpoint protected bisa diakses**

### **Contoh format token di Swagger:**
```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## Test Endpoints

### 1. **Register User (Public)**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com", 
    "password": "password123"
  }'
```

### 2. **Login (Public)**
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

Response:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### 3. **Access Protected Endpoint**
```bash
curl -X GET http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer <access_token>"
```

### 4. **Refresh Token**
```bash
curl -X POST http://localhost:8080/api/v1/refresh-token \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "<refresh_token>"
  }'
```

---

## 📊 **Swagger Features Added:**

✅ **Security Definitions**: Bearer Authentication  
✅ **Protected Endpoints**: Otomatis menampilkan field Authorization  
✅ **API Documentation**: Lengkap dengan contoh request/response  
✅ **Authorization UI**: Tombol "Authorize" untuk input token  

---

## Token Configuration

- **Access Token**: Berlaku 15 menit
- **Refresh Token**: Berlaku 7 hari
- **Secret Key**: `your_secret_key` (ganti di production)

---

## Files Created/Modified:

1. ✅ `pkg/utils/jwt.go` - JWT utilities
2. ✅ `internal/delivery/http/handler/auth_handler.go` - Login & refresh handlers
3. ✅ `internal/delivery/http/middleware/jwt.go` - JWT authentication middleware
4. ✅ `internal/usecase/user_usecase.go` - Added AuthenticateUser method
5. ✅ `internal/domain/entity/user.go` - Added UserLoginRequest struct
6. ✅ `cmd/api/main.go` - Added auth routes, protected endpoints & security definitions
7. ✅ `internal/delivery/http/handler/user_handler.go` - Added @Security annotations

---

## Swagger Documentation

Akses: http://localhost:8080/swagger/index.html

**🎉 Swagger sudah include:**
- Bearer Authentication UI
- Field Authorization untuk endpoint protected
- Dokumentasi lengkap semua endpoint

---

**🎉 Implementasi JWT Authentication selesai dan Swagger sudah support Authorization!**
