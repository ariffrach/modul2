//UNGUIDED 2B KE 3
//IQBAL BAWANI
// 2311102130
///S1IF11D

package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64
	var totalBerat float64
	Iqbalbawani := true

	for Iqbalbawani {
		fmt.Print("Masukkan berat belanjaan di kedua kantong : ")
		fmt.Scan(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Berat kantong tidak boleh negatif!")
			break
		}

		totalBerat = beratKiri + beratKanan
		selisih := math.Abs(beratKiri - beratKanan)

		if selisih > 9 {
			fmt.Printf("Sepeda motor Pak Andi akan oleng: true (Selisih: %.1f)\n", selisih)
			Iqbalbawani = false
			break
		} else {
			fmt.Printf("Sepeda motor Pak Andi akan oleng: false (Selisih: %.1f)\n", selisih)
		}

		if totalBerat > 150 {
			fmt.Printf("Total berat belanjaan melebihi 150 kg: %.1f kg\n", totalBerat)
			Iqbalbawani = false
			break
		}

		if beratKiri > 9 || beratKanan > 9 {
			fmt.Printf("Salah satu kantong melebihi 9 kg! Kantong Kiri: %.1f, Kantong Kanan: %.1f\n", beratKiri, beratKanan)
			Iqbalbawani = false
			break
		}
	}

	fmt.Println("Proses selesai.")
}
