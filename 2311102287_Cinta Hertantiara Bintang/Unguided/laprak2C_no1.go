package main

import (
	"fmt"
)

func main() {
	var (
		total, sisa, diskon, HargaKg int
	)
	fmt.Print("Berat persel (gram) : ")
	fmt.Scanln(&total)
	kg := total / 1000
	gram := total % 1000
	fmt.Printf("Detail berat : %d kg + %d gr", kg,
		gram)
	if total > 10000 {
		HargaKg = kg * 10000
		if gram >= 500 {
			sisa = gram * 5
		} else if gram < 500 {
			sisa = gram * 15
		}
		fmt.Println("\nDetail biaya : Rp.", HargaKg, "+ Rp.", sisa)
		if sisa <= 1000 {
			diskon = 0
		}
		totalbiaya := HargaKg + diskon
		fmt.Println("Total biaya : Rp.", totalbiaya)
	} else {
		HargaKg = kg * 10000
		if gram >= 500 {
			sisa = gram * 5
		} else if gram < 500 {
			sisa = gram * 15
		}
		fmt.Println("\nDetail biaya : Rp.", HargaKg, " + Rp.", sisa)
		totalbiaya := HargaKg + sisa
		fmt.Println("Total biaya : Rp.", totalbiaya)
	}
}
