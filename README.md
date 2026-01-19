# ✨ Bot Keren

**Bot Keren** adalah bot Discord sederhana namun canggih yang dibangun menggunakan **Go (Golang)**. Bot ini terintegrasi langsung dengan kecerdasan buatan **Google Gemini 3 Flash** untuk memberikan respons yang cepat, cerdas, dan natural terhadap pesan pengguna.

## 🚀 Fitur Utama

*   **Integrasi AI Canggih:** Menggunakan model `gemini-3-flash` dari Google untuk percakapan yang luwes.
*   **Respons Otomatis:** Menjawab setiap pesan yang diterima di channel tempat bot berada.
*   **Penanganan Pesan Panjang:** Secara otomatis memecah respons AI yang lebih dari 2000 karakter agar muat di chat Discord.
*   **Typing Indicator:** Menampilkan status "sedang mengetik" saat memproses jawaban AI.
*   **Arsitektur Modular:** Kode terstruktur rapi dengan pemisahan *concern* antara logika bot (`pkg/dcbot`) dan logika AI (`pkg/ai`).

---

## 🛠️ Persiapan

Sebelum menjalankan bot ini, pastikan Anda memiliki:

1.  **Go (Golang):** Terinstal di komputer Anda. [Download di sini](https://go.dev/dl/).
2.  **Discord Bot Token:** Dari [Discord Developer Portal](https://discord.com/developers/applications).
3.  **Google Gemini API Key:** Dari [Google AI Studio](https://aistudio.google.com/).

---

## 📦 Instalasi & Konfigurasi

Ikuti langkah-langkah mudah ini untuk menjalankan Bot Keren di komputer Anda:

### 1. Clone Repository
```bash
git clone https://github.com/username/bot-keren.git
cd bot-keren
```

### 2. Install Dependencies
Download semua library yang dibutuhkan:
```bash
go mod tidy
```

### 3. Konfigurasi Token
⚠️ **PENTING:** Saat ini konfigurasi token masih dilakukan langsung di dalam kode (hardcoded).

*   **Discord Token:**
    Buka file `cmd/main/main.go` dan ganti `TOKEN_BARU_ANDA_DISINI` dengan token bot Discord Anda.
    ```go
    // cmd/main/main.go
    const BotToken = "MASUKKAN_TOKEN_DISCORD_ANDA_DI_SINI"
    ```

*   **Gemini API Key:**
    Buka file `pkg/ai/ai.go` dan ganti `TOKEN_GEMINI_API` dengan API Key Gemini Anda.
    ```go
    // pkg/ai/ai.go
    var GEMINII_API string = "MASUKKAN_API_KEY_GEMINI_ANDA_DI_SINI"
    ```

> 💡 **Tips Keamanan:** Jangan pernah meng-upload file yang berisi token asli ke repository publik (GitHub/GitLab). Untuk produksi, disarankan menggunakan *Environment Variables*.

---

## ▶️ Cara Menjalankan

Setelah konfigurasi selesai, Anda bisa menjalankan bot menggunakan `go run` atau `make`.

### Menggunakan Go Run
```bash
go run cmd/main/main.go
```

### Menggunakan Makefile (Direkomendasikan)
Kami telah menyediakan `Makefile` untuk memudahkan perintah:

*   **Jalankan Bot:**
    ```bash
    make run
    ```
*   **Build Binary:**
    ```bash
    make build
    ```
    Binary akan tersimpan di folder `bin/`.
*   **Bersihkan Project:**
    ```bash
    make clean
    ```

Jika berhasil, Anda akan melihat pesan:
```
Bot sedang berjalan... Tekan CTRL-C untuk berhenti.
```

Sekarang coba chat bot Anda di Discord! 🎉

---

## 📂 Struktur Folder

```
bot-keren/
├── cmd/
│   └── main/
│       └── main.go       # Entry point aplikasi
├── pkg/
│   ├── ai/
│   │   └── ai.go         # Logika interaksi dengan Google Gemini
│   └── dcbot/
│       └── dcbot.go      # Logika bot Discord (handler pesan)
├── go.mod                # Definisi modul Go
└── go.sum                # Checksum dependencies
```

---

## 🤝 Kontribusi

Tertarik mengembangkan bot ini? Silakan fork repository ini dan buat Pull Request! Ide pengembangan selanjutnya:
*   Menambahkan support untuk *Slash Commands*.
*   Menggunakan `.env` untuk menyimpan token agar lebih aman.
*   Menambahkan memori percakapan (context aware).

---

Dibuat dengan ❤️ dan ☕ menggunakan Go.
