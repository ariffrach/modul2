package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Jika salah satu berat negatif, hentikan program
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Berat tidak boleh negatif.")
			break
		}

		// Hitung selisih berat
		selisih := math.Abs(berat1 - berat2)

		// Cek apakah sepeda motor akan oleng
		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}

		// Hentikan program jika salah satu kantong berisi 9 kg atau lebih
		if berat1 >= 9 || berat2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
