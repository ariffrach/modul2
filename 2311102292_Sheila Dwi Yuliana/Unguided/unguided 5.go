package main

import "fmt"

func main() {
	var beratParsel int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scanln(&beratParsel)

	// Hitung berat dalam kg dan sisa gram
	kg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	// Hitung biaya dasar berdasarkan kg
	biayaDasar := kg * 10000

	// Hitung biaya tambahan berdasarkan sisa gram
	var biayaTambahan int
	if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else if sisaGram > 0 {
		biayaTambahan = sisaGram * 15
	}

	// Jika total berat lebih dari 10kg, gratiskan sisa gram
	if kg > 10 && sisaGram > 0 {
		biayaTambahan = 0
	}

	// Hitung total biaya
	totalBiaya := biayaDasar + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gram\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp %d + Rp %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
