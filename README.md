# Aplikasi Autentikasi

Ini adalah aplikasi autentikasi sederhana yang dibuat menggunakan Go dan framework Gin.

## Cara Penggunaan

1. **Masuk**: Kirimkan permintaan POST ke `/login` dengan username dan password untuk mendapatkan token JWT.

2. **Akses Rute Dilindungi**: Gunakan token JWT dalam header `Authorization` untuk mengakses rute yang dilindungi.

## Endpoint

- **POST /login**: Masuk dan dapatkan token JWT.
- **GET /user**: Dapatkan informasi pengguna.

## Kredensial Bawaan

- Username: admin
- Password: 12345

## Otorisasi

Semua pengguna memiliki akses ke semua rute.

