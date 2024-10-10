package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Menghasilkan angka acak antara 1 hingga 100
	rand.Seed(time.Now().UnixNano()) // Mengatur seed menggunakan waktu saat ini
	angkaTebakan := rand.Intn(100) + 1
	berhasil := false

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Tebak angka antara 1 hingga 100. Anda memiliki 5 percobaan.")

	for i := 1; i <= 5; i++ {
		var tebakan int
		fmt.Printf("Percobaan %d: ", i)

		_, err := fmt.Scanln(&tebakan) // Ganti Scan dengan Scanln
		if err != nil {
			fmt.Println("Input tidak valid. Silakan masukkan angka.")
			os.Exit(1)
		}

		// Validasi apakah tebakan di luar batas
		if tebakan < 1 || tebakan > 100 {
			fmt.Println("Tebakan harus antara 1 hingga 100.")
			continue
		}

		// Membandingkan tebakan dengan angka yang benar
		if tebakan == angkaTebakan {
			berhasil = true
			break // Keluar dari loop jika tebakan benar
		} else if tebakan < angkaTebakan {
			fmt.Println("Tebakan Anda terlalu rendah.")
		} else {
			fmt.Println("Tebakan Anda terlalu tinggi.")
		}
	}

	// Menampilkan hasil akhir
	if berhasil {
		fmt.Println("Selamat! Anda berhasil menebak angka dengan benar.")
	} else {
		fmt.Printf("Sayang sekali! Angka yang benar adalah %d.\n", angkaTebakan)
	}
}
