# Golang Hexagonal Architecture
Project Golang dengan implementasi **Hexagonal Architecture**.

## Struktur Direktori
`/cmd` berisi berkas `main.go`, yang digunakan untuk menjalankan aplikasi.

`/internal` berisi semua kebutuhan dari aplikasi.

`/internal/core` berisi logic inti dari aplikasi.

`/internal/core/models` berisi struct atau entity yang dibutuhkan aplikasi untuk mengambil data.

`/internal/core/repository` berisi interface dan implementasinya.

`/internal/core/repository/contract` berisi kontrak interface yang dibutuhkan.

`/internal/core/service` berisi penghubung antara `implementation` dan `controller`.

`/internal/framework` berisi kebutuhan pihak ketiga untuk menjalankan aplikasi.

`/internal/framework/database` berisi komponen penyimpanan yang dibutuhkan aplikasi.

`/internal/framework/routes` berisi jalur *end-point* dari **API**.

`/internal/framework/transport/` berisi controller dan middleware yang digunakan aplikasi.

`/internal/framework/repository/implementation` berisi implementasi dari kontrak interface.

`/internal/framework/repository/mocks` berisi mocking yang dapat digunakan untuk melakukan unit testing dari implementasi kontrak interface.

