// UNGUIDED 2C KE 1
// IQBAL BAWANI
// 2311102130
// S1IF1107

package main

import "fmt"

func main() {
	var berat, kilogram, gram, biayakirim, tambahanbiaya, totalbiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kilogram = berat / 1000
	gram = berat % 1000

	if kilogram > 10 {
		tambahanbiaya = 0
	} else if gram >= 500 {
		tambahanbiaya = gram * 5
	} else {
		tambahanbiaya = gram * 15
	}

	biayakirim = kilogram * 10000
	totalbiaya = biayakirim + tambahanbiaya

	fmt.Printf("Detail berat: %d kg + %d gr\n", kilogram, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayakirim, tambahanbiaya)
	fmt.Printf("Total biaya: Rp. %d\n", totalbiaya)
}
