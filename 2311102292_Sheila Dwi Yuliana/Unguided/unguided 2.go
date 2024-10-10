package main

import "fmt"

func main() {
	var N int
	var bunga string
	var pita string

	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scanln(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		pita += bunga + " - "
	}

	pita = pita[:len(pita)-3]

	fmt.Println("Pita:", pita)
}
