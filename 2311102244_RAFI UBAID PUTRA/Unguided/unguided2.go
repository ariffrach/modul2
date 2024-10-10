package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var pita []string

	fmt.Print("Masukkan jumlah bunga (atau ketik SELESAI untuk berhenti lebih awal): ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)

		if strings.ToLower(bunga) == "selesai" {
			break
		}

		pita = append(pita, bunga)
	}

	hasilPita := strings.Join(pita, " — ")

	fmt.Printf("\nPita: %s\n", hasilPita)
	fmt.Printf("Banyaknya bunga: %d\n", len(pita))
}
