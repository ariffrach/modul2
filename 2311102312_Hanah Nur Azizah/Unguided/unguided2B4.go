package main

import (
	"fmt"
)
//HANAH NUR AZIZAH 2311102312
func hitungAkar2(iterasi int) float64 {
	hasilPerkalian := 1.0 
	
	for indeks := 0; indeks <= iterasi; indeks++ {
		pembilang := (4*float64(indeks) + 2) * (4*float64(indeks) + 2) 
		penyebut := (4*float64(indeks) + 1) * (4*float64(indeks) + 3)  
		hasilPerkalian *= pembilang / penyebut
	}
	return hasilPerkalian 
}

func main() {
	var iterasi int
	fmt.Print("Masukkan jumlah iterasi (K): ")
	fmt.Scan(&iterasi)
	hasilAkar2 := hitungAkar2(iterasi)
	fmt.Printf("Nilai hampiran akar 2 = %.10f\n", hasilAkar2)
}
