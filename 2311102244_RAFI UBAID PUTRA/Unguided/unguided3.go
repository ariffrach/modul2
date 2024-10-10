package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		// Meminta input berat kantong kiri dan kanan
		fmt.Print("Masukkan berat belanjaan di kedua kantong (pisahkan dengan spasi): ")
		fmt.Scanln(&beratKiri, &beratKanan)

		// Cek kondisi apabila salah satu kantong negatif
		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai. Berat kantong tidak boleh negatif.")
			break
		}

		// Cek kondisi apabila total berat melebihi 150 kg
		totalBerat := beratKiri + beratKanan
		if totalBerat > 150 {
			fmt.Println("Proses selesai. Total berat melebihi 150 kg.")
			break
		}

		// Cek apakah selisih berat lebih dari atau sama dengan 9 kg
		selisih := math.Abs(beratKiri - beratKanan)
		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}

		// Cek kondisi apabila salah satu kantong mencapai 9 kg atau lebih
		if beratKiri >= 9 || beratKanan >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
