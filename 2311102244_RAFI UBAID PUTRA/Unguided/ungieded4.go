package main

import (
	"fmt"
)

func main() {
	var beratParsel int
	var biayaPerKg int = 10000
	var biayaTambahanPerGramLebihDari500 int = 5
	var biayaTambahanPerGramKurangDari500 int = 15

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaKg := beratKg * biayaPerKg

	biayaSisaGram := 0
	if beratKg > 10 {
		biayaSisaGram = 0
	} else if sisaGram >= 500 {
		biayaSisaGram = sisaGram * biayaTambahanPerGramLebihDari500
	} else {
		biayaSisaGram = sisaGram * biayaTambahanPerGramKurangDari500
	}

	totalBiaya := biayaKg + biayaSisaGram

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: RP. %d + RP. %d\n", biayaKg, biayaSisaGram)
	fmt.Printf("Total biaya: RP. %d\n", totalBiaya)
}
