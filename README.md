# Pinjemlah

Pinjemlah adalah sebuah aplikasi untuk mengelola pinjaman dan pembayaran. Aplikasi ini dibangun menggunakan Flutter untuk sisi frontend dan Go Fiber untuk sisi backend.

## Fitur

- **Autentikasi Pengguna:** Pengguna dapat mendaftar dan login.
- **Pengelolaan Pinjaman:** Pengguna dapat mengajukan pinjaman, melihat status pinjaman, dan melakukan pembayaran.
- **Admin:** Admin dapat mengelola pengguna, memverifikasi pinjaman, dan melihat laporan.

## Teknologi yang Digunakan

### Frontend
- Flutter
- GetX untuk state management dan routing
- Dio untuk HTTP requests
- TailwindCSS dan DaisyUI untuk styling

### Backend
- Go Fiber
- GORM untuk ORM
- PostgreSQL sebagai database

## Instalasi dan Penggunaan

### Prasyarat

- Flutter SDK
- Go
- PostgreSQL

### Instalasi Backend

1. Clone repository ini:
   ```sh
   git clone https://github.com/nurmanhadi/pinjemlah-fiber.git
