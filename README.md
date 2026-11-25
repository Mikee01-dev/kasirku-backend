# KasirKu — Point of Sales dan Inventory Management API

## Deskripsi Sistem
KasirKu adalah backend REST API untuk kasir dan manajemen stok UMKM. Sistem ini menerapkan keamanan (JWT dan role-based access), integritas data transaksi (database transaction dan row-level locking), serta dokumentasi API (Swagger). Tujuannya meminimalkan salah catat stok dan mencegah stok minus saat banyak transaksi masuk bersamaan.

## Role pada Sistem
| Role | Hak Akses Utama |
| :--- | :--- |
| **Admin** | Mengelola user dan data produk, melihat laporan penjualan |
| **Kasir** | Melihat produk, membuat transaksi (checkout), melihat riwayat |

## Fitur dan Modul Kuliah
| Fitur | Deskripsi Teknis | Modul |
| :--- | :--- | :--- |
| Authentication dan Authorization | Login JWT, password di-hash Bcrypt, role Admin dan Kasir | 5 dan 7 |
| Product Management (CRUD) | Produk, kategori, harga, stok, dengan validasi input | 3 dan 7 |
| Transactions (Checkout) | Checkout multi-item memakai database transaction dan SELECT ... FOR UPDATE | 9 |
| Reporting dan Pagination | Laporan penjualan dengan filter, sort, dan keyset pagination | 3 dan 7 |
| Standardized Error Handling | Response error JSON seragam, tidak membocorkan error database | 7 |
| API Documentation | Dokumentasi interaktif Swagger (OpenAPI) | 12 |

## Rancangan Database (PostgreSQL)
*   **users**: id, name, email, password_hash, role, created_at.
*   **products**: id, sku, name, category, price, stock, is_active, created_at, updated_at.
*   **transactions**: id, user_id, total_amount, payment_method, status, created_at.
*   **transaction_items**: id, transaction_id, product_id, quantity, subtotal.
*   **refresh_tokens**: id, user_id, token_hash, expires_at.

## Struktur Folder
```text
kasirku-backend/
|-- app/
|   |-- model/
|   |-- repository/
|   `-- service/
|-- route/
|-- middleware/
|-- helper/
|-- config/
|-- database/
|-- migrations/
|-- docs/
|-- main.go
|-- .env.example
|-- README.md
```

## Ringkasan Pembagian Tugas (Fase 1)
| Orang | Nama | Branch | File Utama |
| :--- | :--- | :--- | :--- |
| **A (Git Master)** | **Michael** | `main`, `feat/service-product`, `feat/wiring` | `product_service.go`, `main.go`, `README.md` |
| **B** | **Rafael** | `feat/migration`, `feat/database-pkg`, `feat/repository-product` | `001_init.sql`, `postgres.go`, `product_repository.go` |
| **C** | **Sovia** | `feat/model`, `feat/route-product`, `feat/config` | `product.go`, `product_route.go`, `config.go` |

## Cara Menjalankan Proyek
1. Clone repo: `git clone https://github.com/Mikee01-dev/kasirku-backend.git`
2. Masuk folder: `cd kasirku-backend`
3. Copy env: `cp .env.example .env`
4. Sesuaikan isi `.env` dengan database lokal.
5. Jalankan migrasi: `psql -U postgres -d kasirku_db -f migrations/001_init.sql`
6. Jalankan: `go run main.go`