package main

import (
	"fmt"
)
//HANAH NUR AZIZAH 2311102312
func hitungBiayaKirim(beratTotalGram int) int {
	jumlahKg := beratTotalGram / 1000
	sisaGram := beratTotalGram % 1000
	biayaPerKg := 10000
	totalBiayaKg := jumlahKg * biayaPerKg
	biayaTambahan := 0

	if jumlahKg >= 10 {
		biayaTambahan = 0
	} else {
		if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5 
		} else {
			biayaTambahan = sisaGram * 15 
		}
	}

	return totalBiayaKg + biayaTambahan
}

func main() {
	var beratTotalGram int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratTotalGram)
	jumlahKg := beratTotalGram / 1000
	sisaGram := beratTotalGram % 1000
	biayaPerKg := 10000 * jumlahKg
	biayaTambahan := 0

	if jumlahKg >= 10 {
		biayaTambahan = 0 
	} else {
		if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5 
		} else {
			biayaTambahan = sisaGram * 15 
		}
	}

	totalBiaya := biayaPerKg + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", jumlahKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
